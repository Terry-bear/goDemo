package main

import "fmt"

/**
	累加器
 */
func addUpper(str string) func(int) int {
	fmt.Println("str", str)
	n := 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	upper := addUpper("haha")
	fmt.Println(upper(1)) // 11
	fmt.Println(upper(2)) // 13
	fmt.Println(upper(3)) // 16
	fmt.Println(addUpper("hehe")(1))
	fmt.Println(addUpper("nene")(1))
}
