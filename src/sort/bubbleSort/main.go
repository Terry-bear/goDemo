package main

import "fmt"

/**
 冒泡排序
 */
func BubbleSort(arr *[5]int) {
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后:", (*arr))
}

func main() {
	// 定义数组
	arr := [5]int{24, 56, 25, 64, 13}
	BubbleSort(&arr)
}
