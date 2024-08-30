package main

import (
	"fmt"
)

func main() {

	// const 수정불가
	//  var 생략 가능 /변경가능 줄여 서  := 로 대체 가능 > func안에서만 가능하고  var 만 가능
	const name string = "beomsoo"
	const boo bool = true
	// var name1 string = "beom1"
	// name1 = "soo"
	name2 := "test"

	fmt.Println(name2)
}
