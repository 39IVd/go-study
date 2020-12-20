package main

import "testlib"

var i, j int = 1, 2
var b = 1
var c = "hi"

const e int = 10
const (
	Visa   = "Visa"
	Master = "MasterCard"
	Amex   = "American Express"
)

/*
GO Keyword
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
*/
const (
	Apple  = iota // 0
	Grape         // 1
	Orange        // 2
)

func main() {
	println("ho")
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	str := "ABC" // 함수 내에서만 가능
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

	song := testlib.GetMusic("Alicia Keys")
	println(song)
}
