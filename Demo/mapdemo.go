package main

import "fmt"

type TESTDEMO struct {
	i int
}
func main() {
var testMap = make(map[string]int)
	testMap["hehe"] = 2
	fmt.Println(testMap)
}
