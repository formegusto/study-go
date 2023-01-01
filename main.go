package main

import (
	"os"
	"strings"

	"github.com/formegusto/study-go/scrapper"
	"github.com/labstack/echo/v4"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, World !")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("data/" + fileName)

	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)

	// 첨부파일 리턴 기능
	return c.Attachment("data/jobs.csv", "job.csv")
}

func main() {
	e := echo.New()
	
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	e.Logger.Fatal(e.Start(":8080"))
}