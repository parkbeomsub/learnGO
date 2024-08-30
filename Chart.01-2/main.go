package main

import (
	"fmt"
	"strings"
)

// (a int, b int) 를 (a , b int) 로 변경가능 둘다 같은 타입이니깐  리턴타입을 꼭 넣어야함
func multiply(a int, b int) int {
	return a * b

}

func main() {
	fmt.Println(multiply(2, 2))
	// 호출해보기
	fmt.Println(lenAndUpper("beomsoo"))
	// 호출하고 선언으로 사용
	totalLenght, upperName := lenAndUpper("parkbeomsub")
	// 만약 호출을 // 하나만 받고싶으면 _ 로 나타내준다.
	totalLenght1, _ := lenAndUpper("parkbeomsub")
	fmt.Println(totalLenght, upperName)
	fmt.Println(totalLenght1)

	repeatMe("P", "A", "R", "K")
}

// Go에만 있는 기능 리턴을 여러개
func lenAndUpper(name string) (int, string) {

	return len(name), strings.ToUpper(name)

}

// 여러개 인자 받기
func repeatMe(words ...string) {
	fmt.Println(words)
}
