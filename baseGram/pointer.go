package main

import "fmt"

// 交换两个变量的值
func swap_org(a, b *int) {
	*a, *b = *b, *a
}
func swap(a, b int) (int, int) {
	return b, a
}
func main() {
	a, b := 3, 4
	c, d := 4, 5
	fmt.Println("交换前:",a, b)
	swap_org(&a, &b)
	fmt.Println("方法一:交换后:",a, b)
	swap(c, d)
	fmt.Println("方法二:",c, d)
}
