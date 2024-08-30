// map
package main

import "fmt"

func main() {
	maps := map[string]string{"1": "a", "2": "b", "3": "a", "4": "b", "5": "a", "6": "b", "7": "a", "8": "b"}
	for key, val := range maps {
		fmt.Println(key, val)
	}

}
