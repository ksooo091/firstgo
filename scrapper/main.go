package main

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

var keyWord string = "데브옵스"
var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + keyWord

type extractedJob struct {
	id       string
	title    string
	location string
}

func main() {
	var jobs []extractedJob
	totalPages := getTotalPages()
	fmt.Println(totalPages)
	for i := 0; i < totalPages+1; i++ {
		extractedJobs := getPage(i + 1)
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)

}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkError(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"id", "Title", "Location"}

	wErr := w.Write(headers)
	checkError(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?view_type=search&rec_idx=" + job.id, job.title, job.location}
		jwErr := w.Write(jobSlice)
		checkError(jwErr)
	}
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
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

		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("value")
	title := cleanString(card.Find(".area_job>.job_tit>a").Text())

	location := cleanString(card.Find(".area_job>.job_condition>span>a").Text())

	return extractedJob{
		id:       id,
		title:    title,
		location: location}
}

func getTotalPages() int {
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

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
