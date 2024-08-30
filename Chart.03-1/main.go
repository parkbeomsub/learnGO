// goroutines
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"beomsoo", "euna"}
	for _, person := range people {

		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for ", i)

		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {

	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- person + " is sexy"
}
