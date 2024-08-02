// array and slices
package main

import "fmt"

func main() {
	//array는 길이가 있는 배열
	name := [5]string{"a", "b", "c"}
	//slices 는 길이를 모를때 무한정할때
	name2 := []string{"a", "b", "c"}
	name2 = append(name2, "appended things")
	fmt.Println(name)
	fmt.Println(name2)

}
