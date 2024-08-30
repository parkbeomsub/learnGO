// struct
package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	fF := []string{"bab", "ramen"}
	beomsoo := person{name: "beomsoo", age: 32, favFood: fF}
	fmt.Println(beomsoo)
}
