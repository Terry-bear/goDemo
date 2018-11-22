package main

import (
	"queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	fmt.Println(q)
	q.Push(2)
	q.Push(3)
	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q.IsEmpty())
}
