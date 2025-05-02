package main

import (
	"fmt"
	"sync"
	"time"
)

func testWaitGroup(v int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函數結束時執行
	fmt.Println("testWaitGroup", v)
	time.Sleep(time.Second * 2) // 模擬耗時操作
}

func testWaitGroupWithChannel(v int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函數結束時執行
	fmt.Println("testWaitGroup", v)
	time.Sleep(time.Second * 2) // 模擬耗時操作
	ch <- v
}
