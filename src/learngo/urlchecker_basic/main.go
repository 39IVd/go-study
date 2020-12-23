package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	// var results = map[string]string
	/* => 선언만 되고 초기화되지 않음
	초기화되지 않은 map에 값을 할당하면 panic 발생
	Panic : 컴파일러가 캐치하지 못하는 에러 */

	// solution 1 : map을 초기화
	// var results = map[string]string{}
	// solution 2 : make() 사용
	// make() : map을 만드는 함수
	var results = make(map[string]string)

	urls := []string{"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}
func hitURL(url string) error {
	fmt.Println("Checking :", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		/* StatusCode : 100, 200, 300 -> redirect
		400이상 -> 에러  */
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
