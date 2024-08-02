package main

// 리턴값 정의와 defer
import (
	"fmt"
	"strings"
)

// 리턴값을 여러 인자를 두는 방법  리턴에 변수와 타입을 지정가능  생략도 가능하고 return에 기입해도 된다.
// defer는 함수가 끝날때  실행된다.
func lenAadUpper(name string) (lenght int, uppercase string) {

	lenght = len(name)
	uppercase = strings.ToUpper(name)

	defer fmt.Println("lenAadUpper func is finished")
	return

}

// 호출은 변수 선언과 같이 :=으로
func main() {

	totalLenght, uppercase := lenAadUpper("base-on")
	fmt.Println(totalLenght)
	fmt.Println(uppercase)

}
