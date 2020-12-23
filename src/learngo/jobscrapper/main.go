package main

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

/*
NOTE:
goquery : html 내부를 볼 수 있는 라이브러리
https://github.com/PuerkitoBio/goquery
1. go get github.com/PuerkitoBio/goquery
2. package import : "github.com/PuerkitoBio/goquery"
3.

- 과정
1. 각 페이지 방문
2. 페이지로부터 job 추출
3. 추출한 job들을 엑셀에 저장

NOTE:
- 문법
. : class 명
> : element안의 element를 참조 ex) .title>a : 클래스명이 title인 element 안의 a 태그를 가진 element
Find() : 특정 Selection안에서 selector string을 포함하는 모든 descendants를 반환
 ex) *Selection.Find(selector)
Each() : 각각의 element에 대해 동작 수행
Attr(attrname) : Selection에서 attrname 속성의 값을 반환
*/

// http://www.jobkorea.co.kr/Search/?stext=golang&tabType=recruit&Page_No=2
// tplPagination newVer wide

type extractedJob struct {
	id       string
	title    string
	company  string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python"

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	// fmt.Println(totalPages)
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		// getPage()는 []extractedJob을 반환한다.
		// append(targetSlice, elems...) 를 통해,
		// 여러 개의 elems slice들을 전부 targetSlice에 추가할 수 있다. => 하나의 slice로 만듬
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))

}

func writeJobs(jobs []extractedJob) {
	// file 생성
	file, err := os.Create("jobs.csv")
	checkErr(err)
	// fileWriter 생성 (file에 write할 수 있는 객체)
	w := csv.NewWriter(file)
	// 함수가 끝나는 시점에 data write
	defer w.Flush()
	headers := []string{"Link", "Title", "Company", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{
			"https://kr.indeed.com/viewjob?jk=" + job.id,
			job.title,
			job.company,
			job.location,
			job.salary,
			job.summary,
		}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}

}

// 2. getPage : 각 Page마다 http.Get() -> goquery로 res.Body의 Document를 읽어 각 job card를 searchCards에 저장
/* indeed 사이트의 page 동작 방식 :
각 페이지마다 50개씩 출력됨 (limit=50)
1페이지는 start=0
limit=50&start=page*50 */
func getPage(page int) []extractedJob {
	var jobs []extractedJob // 한 페이지의 job slice
	pageURL := baseURL + "&limit=50" + "&start=" + strconv.Itoa(page*50)
	/* page*50 : int이므로, string에 합치려면 string으로 변환해야 함!
	GO의 내장 함수인 strconv.Itoa() 사용 : string conversion Integer to ASCII */
	fmt.Println("Requesting", pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		// 각 카드에 대해,
		// 1. extractJob()으로 job 정보를 추출해 extractedJob struc에 저장하고,
		// 2. extractedJob slice인 jobs에 해당 job을 추가한다.
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs

}

// 3. extractJob : searchCards안의 모든 card로부터 job slices 추출 후 저장
func extractJob(card *goquery.Selection) extractedJob {
	// card : jobsearch-SerpJobCard라는 클래스명을 가진 element
	// card의 실제 값 : <div class="~~~" data-jk="~~~"..>
	// 이 때, div, data-jk는 모두 s의 Attribute(속성)이다.

	// card에서 data-jk라는 attr를 가져옴 : string, exist를 반환
	id, _ := card.Attr("data-jk")

	// 클래스명이 title인 element안의 a 태그 element를 찾아서 text 추출
	title := cleanString(card.Find(".title>a").Text())
	company := cleanString(card.Find(".company").Text())
	location := cleanString(card.Find(".location").Text())
	// location := card.Find(".sjcl").Find(".location").Text()
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	return extractedJob{
		id:       id,
		title:    title,
		company:  company,
		location: location,
		salary:   salary,
		summary:  summary,
	}
	// fmt.Println(id, title, company, location, salary, summary)

}

// jobsearch-SerpJobCard

// 1. getPages() : 총 Page 수를 반환하는 함수
func getPages() int {
	pages := 0
	// url에 접속 : http Request
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)
	// response Body : byte type이므로 함수가 끝나면 닫아줘야 한다. => memory leak 방지
	defer res.Body.Close()

	// goquery document 반환 : HTML document를 load
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// Find()로 document안의 클래스명이 pagination인 태그를 찾고,
	// Each()로 각각의 태그에 대해 동작 수행
	// indeed 사이트에는 pagination이 하나이다.
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// => Selection 객체를 반환
		// fmt.Println(s.Html()) // 각 Selection의 HTML 반환

		pages = s.Find("a").Length() // Selection에서 a 태그들의 개수 반환
	})
	return pages
}
func checkErr(err error) {
	// error가 존재하면 프로그램 종료
	if err != nil {
		log.Fatalln(err)
	}
}
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed withs Status : ", res.StatusCode)
	}
}
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
	/*
		NOTE: strings 라이브러리
		- strings.TrimSpace() : string 양 끝의 공백 제거
		- strings.Fields() : 모든 공백을 제거하고 string array 반환
		- strings.Join([]string, seperator) : string 배열을 가져와서 사이사이에 seperator를 넣고 합친 결과 반환
		ex)
		"hello   f    1"
		Fields => ["hello","f","1"]
		Join([], " ") => "hello f 1"
	*/
}
