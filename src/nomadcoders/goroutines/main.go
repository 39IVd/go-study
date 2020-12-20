package main

import (
	"fmt"
	"time"
)

/*
NOTE:
Goroutines : 다른 함수와 동시에 실행시키는 함수
함수명 앞에 go 만 붙여도 다른 함수와 동시에 실행 가능
!! Goroutines은 main 함수가 종료되면 소멸된다.
(main 함수는 다른 goroutines를 기다리지 않는다.)
*/
func main() {
	go sexyCount("paige")
	// go sexyCount("lee")
	// => 프로그램이 바로 종료됨 : main함수가 바로 끝나버렸기 때문
	sexyCount("lee")
	// => 이 함수가 실행되는 동안에는 main이 종료되지 않기 때문에 유효하다.

}
func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
