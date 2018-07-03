package main

import "fmt"

/**
	累加器
 */
func addUpper() func(int) int {
	n := 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	upper := addUpper()
	fmt.Println(upper(1)) // 11
	fmt.Println(upper(2)) // 13
	fmt.Println(upper(3)) // 16
	fmt.Println(addUpper()(1))
	fmt.Println(addUpper()(1))
}
