package main

import (
	"github.com/firstgo/echo/scrapper"
	"github.com/labstack/echo"
	"os"
	"strings"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("echo/index.html")
}

func handleScrape(c echo.Context) error {
	defer func() {
		err := os.Remove(fileName)
		if err != nil {

		}
	}()
	keyWord := strings.ToLower(scrapper.CleanString(c.FormValue("keyWord")))
	scrapper.Scrape(keyWord)
	return c.Attachment(fileName, "job.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))

}
