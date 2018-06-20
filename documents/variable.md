##  fmt.Printf() 可以用于做格式化输出
```
fmt.Printf('n1 的类型是 %T', n1)
```

> %T 获取当前变量的类型

## 如何在程序中查看某个变量的占用字节大小和数据类型
```
var n2 int64 = 100
fmt.Printf('n2 的类型%T \n n2占用的字节数是%d', n2, unsafe.Sizeof(n2))
```

> %d 获取字节大小

## 当我们直接输出byte值,就是输出了对应的字符码值
```
var c1 byte = 'a'
var c2 byte = '0'

// 直接输出码值
fmt.Println("c1=", c1)
fmt.Println("c2=", c2)

// 输出对应的字符
fmt.Println("c1=%c, c2=%c", c1, c2)
```

