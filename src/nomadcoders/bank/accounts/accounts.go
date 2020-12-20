package accounts

import (
	"errors"
	"fmt"
)

// Account struct : struct name, field 모두 대문자여야만 public으로 간주 -> 외부에서 접근 가능
// 소문자 : private, 대문자 : Public
type Account struct {
	owner string
	// Balance int
	balance int
}

// error를 변수에 할당
var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
/*
외부에서 struct에 접근 가능하지만 값의 변경을 막고 싶을 경우 : 생성자 구현 (constructor)
*/
func NewAccount(owner string) *Account {
	/*
		*Account : Account struct의 복사본을 return
		=> Accout의 복사본을 만들고 싶지만, 새로운 object를 만들긴 싫은 경우
		: 복사본 자체를 return
	*/
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
/* NOTE:
1. 함수 (Function)
func 함수명(변수 type(argument)) 리턴타입

2. 메소드 (Method)
func (변수 type(struct reciever)) 함수명(변수 type(argument)) 리턴타입
=> struct가 해당 method를 가짐
: GO에서는 struct안에 method가 존재하는 것이 아니라,
functon을 만들어 function name 앞에 (reciever struct)를 명시하면
자동적으로 struct가 해당 method를 가지게 됨
*/
// func (a Accout) Deposit(amount int) {
// => Account a reciever object를 복사
func (a *Account) Deposit(amount int) {
	/*
		(a *Accout) : pointer reciever
		=> Account의 복사본을 만들지 않고, Deposit()을 호출한 account를 그대로 사용

		(a Account) a : reciever, Account : type
		reciever : struct명의 첫 글자를 따서 소문자로 지어야 함!
	*/
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw from your account
/*
NOTE: Error Handling
GO에서는 try catch, exception과 같은 에러 처리 함수가 존재하지 않음
=> error를 직접 체크하고 return해야 함
: error에는 (error, nill)의 두 가지 type이 존재함
*/
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// return errors.New("Can't withdraw")
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Override String() method
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas : ", a.Balance())
}
