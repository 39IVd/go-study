package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestresult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request Failed")

func main() {
	results := make(map[string]string)
	c := make(chan requestresult)

	urls := []string{"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/"}
	/*
		NOTE:
		동시성 처리가 필요한 이유 :
		- 코드가 순서대로 주어져 있고, 모두 1초만에 처리가 끝나는데 마지막 코드는 2초가 걸린다고 가정한다.
		- 프로그램이 순차적으로 처리되면 총 실행시간은 전부 다 합친 시간이 걸릴 것이다.
		- 동시성 처리를 통해 모든 코드가 동시에 동작하도록 하면, 프로그램은 최대 2초 걸릴 것이다.
	*/
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		// fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}
	// => Goroutines을 사용하면 총 실행시간이 가장 오래걸리는 url 하나의 시간과 동일해진다!!
	for url, status := range results {
		fmt.Println(url, status)
	}

}
func hitURL(url string, c chan<- requestresult) {
	// NOTE: chan<- : Send Only Channel (전송만 가능한 채널)
	// fmt.Println(<-c) // Send Only이므로 수신 불가
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestresult{url: url, status: status}

}
