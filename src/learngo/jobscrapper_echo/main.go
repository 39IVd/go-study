package main

import (
	"fmt"
	"learngo/jobscrapper_echo/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

// go get github.com/labstack/echo

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}
func handleScrape(c echo.Context) error {
	// 함수가 끝나면 서버의 파일을 삭제
	defer os.Remove(fileName)
	fmt.Println(c.FormValue("term"))
	// html의 <form>안에 "term"이라는 이름을 가진 element가 반환하는 값 리턴
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	// c.Attachment(a, b) : 첨부파일을 리턴하는 기능
	// a : 반환할 파일 이름
	// b : 사용자에게 전달할 파일 이름
	return c.Attachment(fileName, fileName)
}
func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	// Port Number : 1323
	e.Logger.Fatal(e.Start(":1323"))
	// scrapper.Scrape("term")
}
