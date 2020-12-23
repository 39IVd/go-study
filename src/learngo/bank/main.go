package main

import (
	"fmt"
	"learngo/bank/accounts"
)

func main() {
	account := accounts.NewAccount("paige")
	fmt.Println(account) // => struct가 자동으로 String() method를 호출함
	// result : &{paige 0} => 복사본이 아니라 object라는 의미
	account.Deposit(10)
	fmt.Println(account.Balance())

	// NOTE: Error Handling
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err) // error를 print하고 프로그램을 종료시킴
		fmt.Println(err)
	}
	account.ChangeOwner("lee")
	fmt.Println(account.Balance(), account.Owner())
}
