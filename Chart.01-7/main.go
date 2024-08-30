// Pointer
package main

import "fmt"

// 주소값 &   인자값 *
func main() {
	a := 3
	b := &a
	*b = 5
	fmt.Println(*b)

}
