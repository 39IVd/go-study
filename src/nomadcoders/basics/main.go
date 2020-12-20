/*
GO Tutorial
*/
// main.go 파일 : 컴파일하고자 하는 파일
// => entry point이므로 가장 먼저 찾아내는 파일

package main

import (
	"fmt"
	"strings"
)

// 2. function
func multiply(a, b int) int {
	return a * b
}
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
func lenAndUpper2(name string) (length int, uppercase string) {
	// return value를 명시
	defer fmt.Println("I'm done")
	// 3. defer : function이 끝난 후에 실행되는 코드!!
	length = len(name)
	uppercase = strings.ToUpper(name)
	return // naked return
}
func repeatMe(words ...string) {
	fmt.Println(words)
}

// 3. loop
func superAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}
	// method 2
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	return total
}

// 4. condition : if, switch
func canIDrink(age int) bool {
	// variable expression : 조건을 쓰는 순간에 variable 생성 가능
	// => if else 문에서만 사용할 변수 선언
	// ; 이전 => 변수 생성
	// ; 이후 => 조건 지정
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
func canIDrink2(age int) bool {
	// method 1
	// switch age {
	// case 18:
	// 	return true
	// }
	switch koreanAge := age + 2; {
	// variable expression 사용 가능
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	case koreanAge > 50:
		return false
	}
	return false
}

// 8. Structs : go에서는 class, object, constructor(생성자)가 존재하지 않음
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// 1. type
	var name string = "s"
	name = "b"
	name2 := false // := 축약형 : 변수만 가능, 함수 내부에서만 가능
	fmt.Println(name, name2)

	// 2. function
	fmt.Println(multiply(2, 2))
	totalLen, _ := lenAndUpper("paige")
	// _ : 컴파일러가 무시하는 underscore
	fmt.Println(totalLen)
	fmt.Println(lenAndUpper2("lee"))
	repeatMe("pai", "ri", "lee")

	// 3. loop : for 밖에 없음
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	// 4. condition : if, switch
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrink2(18))

	// 5. pointer
	// pointer : 값을 복사해서 저장하지 않고, 같은 객체를 같은 주소를 참조하여 여러 번 재사용할 때 사용 => 성능 개선 시 필요
	a := 2
	// b := a // copy value (값 복사)
	b := &a     // a의 메모리 주소를 b에 복사 : copy reference (주소 복사)
	a = 10      // a를 10으로 변경 => b의 실제 값도 10으로 변경됨
	*b = 202020 // b의 실제값을 20으로 변경 => a 값도 변경됨
	// (값이 변경되어도 주소는 그대로임)
	fmt.Println(&a, b)
	fmt.Println(a, *b) // &a : 메모리 주소, *b : b에 저장된 실제 값

	// 6. Arrays and Slices
	// array
	names := [5]string{"paige", "lee", "joo"}
	names[0] = "bb"
	// names[6] = "aa" // 에러 발생 (메모리 초과)
	fmt.Println(names)
	// slices (array와 비슷하지만 length가 없음)
	names2 := []string{"paige", "lee", "joo"}
	names2 = append(names2, "ll")
	// append : slice에 값을 추가해 새 slice return
	fmt.Println(names2)

	// 7. Maps : map[key]value
	lee := map[string]string{"name": "lee", "age": "13"}
	fmt.Println(lee)
	for key, value := range lee {
		fmt.Println(key, value)
	}

	// 8. Structs
	favFood := []string{"apple", "pear"}
	// paige := person{"paige", 18, favFood}
	paige := person{name: "paige", age: 18, favFood: favFood}
	fmt.Println(paige.name, paige.age, paige.favFood)
}
