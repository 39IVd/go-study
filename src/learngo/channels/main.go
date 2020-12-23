package main

import (
	"fmt"
	"time"
)

/*
NOTE:
Channel : Goroutines과 main()함수 간의 정보를 전달하기 위한 방법
(Goroutines간의 정보 전달도 가능 - Communication)
- Channel로부터 결과를 받을 때, main()함수는 결과값이 반환될 때까지 기다린다.
- Goroutines는 return value를 반환할 수 없음
1. 특정 type의 value를 보낼 수 있는 Channel을 생성하고,
2. Goroutines 함수의 argument에 해당 Channel을 할당한다.
3. Goroutines 함수에서 저장할 값을 <- 를 이용해 저장한다.
4. main()함수에서 Channel에 저장된 값을 기다렸다가 받는다.
- ' <- ' : blocking operation (Channel로부터 메세지를 기다렸다가 가져옴)

NOTE: Channel Rule
- 먼저 실행되는 Goroutines을 먼저 Channel에 저장한다.
- main() 함수가 종료되면 Goroutines는 소멸된다.
- Channel을 통해 보내고 받을 data type을 구체적으로 명시해야 한다.
- 메세지 전송 : chan <- 값
- 메세지 수신 : 값 <- chan or <- chan

NOTE: Blocking Operation
- 프로그램 (이 경우엔 main())이 Channel의 값을 받기 전까지 동작을 멈춤

*/
func main() {
	c := make(chan bool) // bool값을 보낼 수 있는 channel 생성
	str := make(chan string)
	// people := [2]string{"paige", "lee"}
	people := []string{"paige", "lee", "joo"}
	for _, person := range people {
		// result := go isSexy(person) // 불가능 !
		go isSexy(person, c, str)
	}
	/* NOTE:
	for문을 통해 Goroutines 2개가 동시에 실행된다 !! (순서 없음)
	Concurrency (동시성)로 인해 결과 출력 순서는 뒤죽박죽임 */

	/*
		// method 1
		result := <-c
		fmt.Println(result, <-str)
		fmt.Println(<-c, <-str)
		// => 먼저 실행이 끝난 Goroutines의 Channel이 먼저 반환됨
		// fmt.Println(<-c)
	*/
	/* NOTE:
	=> Channel을 하나 더 받아오면 Deadlock 발생 !
	: 실제로 만든 Goroutines는 2개의 Channel을 반환하는데,
	하나의 Channel을 더 기다리고 있기 때문 */

	// method 2
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for", i)
		// blocking operation이므로, <- chan의 값이 받아오기 전까지 동작을 멈춤
		fmt.Println(<-c, <-str)

	}
}
func isSexy(person string, c chan bool, str chan string) {
	time.Sleep(time.Second * 3)
	fmt.Println(person)
	c <- true
	str <- person + " is sexy"
}
