# Golang 中函数的闭包
## 基本介绍
> 闭包就是一个函数和与其相关的引用环境组合的一个整体

```go
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
```
