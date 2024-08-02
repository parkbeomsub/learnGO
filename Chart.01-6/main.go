// d
package main

import "fmt"

func canIDrink(age int) bool {

	// 	// 아래와 같이 조건문에  초기 내용 변경 가능
	// 	if koreaAge := age; koreaAge < 18 {
	// 		return false
	// 	}
	// 	return true

	// }

	switch koreaAge := age + 2; koreaAge {
	case 18:
		return false
	case 20:
		return true
	case 32:
		return false

	}
	return false

}

func main() {

	fmt.Println(canIDrink((32)))

}
