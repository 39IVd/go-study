package main

import (
	"fmt"
	"learngo/dictionary/mydict"
)

func main() {
	word := "hello"
	definition := "greeting"
	dictionary := mydict.Dictionary{}
	// dictionary["hello"] = "hello"
	// fmt.Println(dictionary)
	// definition, err := dictionary.Search("second")
	dictionary.Add(word, definition)
	dictionary.Search(word)

	err2 := dictionary.Add(word, "hi")
	if err2 != nil {
		fmt.Println(err2)
	}
	err3 := dictionary.Update(word, "hi")
	if err3 != nil {
		fmt.Println(err3)
	}
	dictionary.Delete(word)
	def2, err4 := dictionary.Search(word)
	print(def2)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(def2)
	}
}
