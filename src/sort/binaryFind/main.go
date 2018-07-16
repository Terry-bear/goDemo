package main

import "fmt"

/**
二分查找法
 */
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	// 判断边界
	if leftIndex > rightIndex {
		fmt.Println("can't find val")
		return
	}

	// 先找到下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		// 要查找的val, 在leftIndex----middle - 1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		// 要查找的val, 在middle + 1 ---- rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到下标为%v \n", middle)
	}
}
func main() {
	arr := [6]int{2, 4, 21, 12, 15, 10}
	BinaryFind(&arr, 0, len(arr)-1, 21)
}
