package scrapper

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type extractedJob struct {
	id       string
	title    string
	location string
}

func Scrape(keyWord string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + keyWord

	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getTotalPages(baseURL)
	fmt.Println(totalPages)
	for i := 0; i < totalPages+1; i++ {
		go getPage(i+1, baseURL, c)

	}
	for i := 0; i < totalPages+1; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)

}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkError(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"링크", "제목", "위치"}

	wErr := w.Write(headers)
	checkError(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?view_type=search&rec_idx=" + job.id, job.title, job.location}
		jwErr := w.Write(jobSlice)
		checkError(jwErr)
	}
}

func getPage(page int, baseURL string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)

	res, err := http.Get(pageURL)
	checkStatusCode(res)
	checkError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)
	//
	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) {

		go extractJob(card, c)

	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := CleanString(card.Find(".area_job>.job_tit>a").Text())

	location := CleanString(card.Find(".area_job>.job_condition>span>a").Text())

	c <- extractedJob{
		id:       id,
		title:    title,
		location: location}
}

func getTotalPages(baseURL string) int {
	pages := 0
	resp, err := http.Get(baseURL)
	checkError(err)
	checkStatusCode(resp)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	doc, err2 := goquery.NewDocumentFromReader(resp.Body)
	checkError(err2)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}
func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkStatusCode(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request Fail!! :", resp.Status)
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
