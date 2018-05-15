package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 5, 6, 8}
	var grid [4][5]bool
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 通过使用_,来省略变量,golang中定义的变量必须用,否则用_标识
	for i, v := range arr3 {
		fmt.Println("index:",i,"item:", v)
	}

}
