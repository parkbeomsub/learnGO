package main

import (
	"fmt"
	"mydict/mydict"
)

func main() {

	// dictionary := mydict.Dictionary{"frist": "first word"}
	// fmt.Println(dictionary["frist"])

	// dictionary["hello"] = "hello"
	// fmt.Println(dictionary)
	// fmt.Println("----")

	// definiton, err := dictionary.Search("Second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definiton)
	// }

	// err = dictionary.Add("hello1", "hello2")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, err := dictionary.Search("hello2")

	// fmt.Println(hello)

	dictionary := mydict.Dictionary{}
	word := "hello"
	dictionary.Add(word, "bye")
	err := dictionary.Update(word, "what's up?")
	if err != nil {
		fmt.Println(err)
	}
	word, _ = dictionary.Search((word))
	fmt.Println(word)

}
