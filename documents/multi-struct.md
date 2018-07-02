# 函数自身作为数据类型
## 函数可以是数据类型进行传递

```go
func add (a int, b int) int {
    return a + b
    }

func multiAdd (funcStru func (n1 int, n2 int) int, a int, b int) int {
    return funcStru(a, b)
    }

func main () {
    res := multiAdd(add, 10, 20)
    fmt.println("res=%d", res)
    }
```

## go支持自定义数据类型
### 基本语法
    type 自定义数据类型名 数据类型
    > eg: type myInt int   // 这时myInt就等价于int
    > eg: type myFunc func(int, int) int

### 支持对函数返回值命名
```go
func cal (n1 int, n2 int)(sum int, sub int){
    sum = n1 + n2
    sub = n1 - n2
    return
    }
```