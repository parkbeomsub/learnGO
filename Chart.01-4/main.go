// 반복문

package main

import "fmt"

func superAdd(numbers ...int) int {

	// 1. range 로  프린트하면 첫번째 값에 인덱스 두번째에는 배열의 값이 들어간다.
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	// 2. 아래 내용으로 선언이 가능하며 c언어와 비슷하게 사용가능
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {

	total := superAdd(1, 2, 2, 2, 2, 6, 7)
	fmt.Println(total)

}
