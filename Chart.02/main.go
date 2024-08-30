package main

import (
	"accounts/accounts"
	"fmt"
)

func main() {
	user1 := accounts.NewCreate("beomsoo")
	user1.Deposit(10)
	err := user1.Withdraw(11)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user1)

}
