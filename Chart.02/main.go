package main

import (
	"fmt"

	"github.com/serranoarevalo/learngo/banking"
)

func main() {

	account := banking.Account{Owner: "beomsoo", Balance: 20000}
	fmt.Println(account)

}
