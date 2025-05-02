package main

import "time"

func channelSender1(v int, c chan int) {
	// for i := 0; i < v; i++ {
	// 	c <- i
	// }
	// close(c) // 關閉通道
	time.Sleep(time.Second * 1)
	c <- 111
}

func channelSender2(v int, c chan int) {
	time.Sleep(time.Second * 2)
	c <- 222
}
