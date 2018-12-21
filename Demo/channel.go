package main

import "fmt"

/**
	todo 通道(Channels) 是连接多个 Go 协程的管道。你可以从一个 Go 协程 将值发送到通道，然后在别的 Go 协程中接收。
 */
func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg + "\n")

	noCacheChannelDemo()
}

/**
	todo nocache channel
	default channel has not cache , if you want add cache ,should make direct cache number
	因为这个通道是有缓冲区的，即使没有一个对应的并发接收 方，我们仍然可以发送这些值。
 */
func noCacheChannelDemo()  {
	noBuf := make(chan string)
	noBuf <- "noBuffered"
	fmt.Println(<-noBuf)
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}