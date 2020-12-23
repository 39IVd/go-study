package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	company  string
	location string
	salary   string
	summary  string
}

// Scrape indeed by a term
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term
	var jobs []extractedJob
	mainC := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(i, mainC, baseURL)
	}
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-mainC
		jobs = append(jobs, extractedJobs...)
	}

	// csv 파일 생성 후 저장
	writeC := make(chan []string)
	file, err := os.Create("jobs.csv")
	checkErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()
	headers := []string{"Link", "Title", "Company", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)
	for _, job := range jobs {
		go writeJobs(job, writeC)
	}
	fmt.Println("Done, extracted", len(jobs))
	for i := 0; i < len(jobs); i++ {
		jwErr := w.Write(<-writeC)
		checkErr(jwErr)
	}
}

/*
NOTE:
Goroutines, Channel을 이용한 Job Scrapper
각 페이지에는 50개의 일자리 정보가 있고, 페이지는 5개이므로 50*5개의 Goroutines이 생성됨
+ 총 job 수는 50*5개이므로 250개의 writeJobs Goroutines이 생성됨 = 500개
- 3개의 Channel 생성
1. main <-> getPage
2. getPage <-> extractJob
3. main <-> writeJobs
*/

// 1. getPages() : 총 Page 수를 반환하는 함수
func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {

		pages = s.Find("a").Length()
	})
	return pages
}

// 2. getPage : 각 Page마다 http.Get() -> goquery로 res.Body의 Document를 읽어 각 job card를 searchCards에 저장
func getPage(page int, mainC chan<- []extractedJob, url string) {
	var jobs []extractedJob
	// getPage <-> extractJob 간 통신 가능한 Channel 생성
	c := make(chan extractedJob)
	pageURL := url + "&limit=50" + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	// 생성된 Goroutines 개수만큼 channel 값 받아옴
	for i := 0; i < searchCards.Length(); i++ {
		// 메세지가 전달되기를 기다렸다가, 전달받으면 jobs 배열에 추가
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

// 3. extractJob : searchCards안의 모든 card로부터 job slices 추출 후 저장
// c chan<- extractedJob : Send-Only channel (channel에 값 전송)
func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	company := CleanString(card.Find(".company").Text())
	location := CleanString(card.Find(".location").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())
	// channel에 값 전송
	c <- extractedJob{
		id:       id,
		title:    title,
		company:  company,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

// 4. writeJobs
func writeJobs(job extractedJob, writeC chan<- []string) {
	writeC <- []string{
		"https://kr.indeed.com/viewjob?jk=" + job.id,
		job.title,
		job.company,
		job.location,
		job.salary,
		job.summary,
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed withs Status : ", res.StatusCode)
	}
}
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
