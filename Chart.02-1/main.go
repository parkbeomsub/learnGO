package main

import (
	"fmt"
	"mydict/mydict"
)

func main() {

	dictionary := mydict.Dictionary{"frist": "first word"}
	fmt.Println(dictionary["frist"])

	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
	fmt.Println("----")

	definiton, err := dictionary.Search("Second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definiton)
	}
}
