package main

import (
	"fmt"
	"time"
)

func pinger(c chan int) {
	for i := 0; ; i++ {
		fmt.Println(6)
		c <- 5
	}
}
func printer(c chan int) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
		fmt.Println(4)
	}
}

func main() {
	var c = make(chan int)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}

// When pinger attempts to send a message on the channel
// it will wait until printer is ready to receive the message.
// (this is known as blocking)
// printer 执行到 `msg := <-c` 的时候，pinger才能发送消息进 channal 类似中断
// 665 456 456 456
// block in 6
