# goLang中的slice
## slice底层实现
* slice是一个引用类型
* 从底层来说,slice是一个数据结构(struct结构体)
```go
type slice struct{
  prt *[2]int
  let int
  cap int
}
```

## slice使用的三种方式
### 方式一:定义一个切片,然后让切片去引用一个已经创建好的数组
```go
var arr[3] int = [...]int{1, 3, 4}
var slice int = arr[0:1]
```
### 方式二:通过make来创建切片
>基本语法: var 切片名 []type = make([]type,len,[cap])
* 切片：size指定了其长度。该切片的容量等于其长度。切片支持第二个整数实参可用来指定不同的容量；
     它必须不小于其长度，因此 make([]int, 0, 10) 会分配一个长度为0，容量为10的切片。
* 映射：初始分配的创建取决于size，但产生的映射长度为0。size可以省略，这种情况下就会分配一个
     小的起始大小。
* 通道：通道的缓存根据指定的缓存容量初始化。若 size为零或被省略，该信道即为无缓存的。

```go
var slice []int = make([]int, 4)

```
### 方式三:定义一个切片后,直接指定具体数组(常用)
```go
var slice []int = []int {1,2,3,4}
```
