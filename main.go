package main

import (
	"fmt"
	"runtime"
	"sync"
)

var globalVar int = 3333

const ccc int = 99

func main() {
	// testGeneric(123)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("發生錯誤：", err)
		}
	}()

	// c := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ { // 多次寫入通道
	// 		c <- i
	// 	}
	// 	close(c) // 關閉通道否則會造成死鎖
	// }()

	// fmt.Println("value", <-c) // 0
	// fmt.Println("value", <-c) // 1

	// for val := range c { // 通道阻塞，直到有資料可以接收 連續接收資料直到 channel 關閉
	// 	fmt.Println("接收資料：", val)
	// }

	// c1 := make(chan int)
	// c2 := make(chan int)
	// go channelSender1(10, c1)
	// go channelSender2(10, c2)

	// select { // 多路複用
	// case msg1 := <-c1:
	// 	fmt.Println("Received from c1", msg1)
	// case msg2 := <-c2:
	// 	fmt.Println("Received from c2", msg2)
	// }

	fmt.Println("起始 goroutines:", runtime.NumGoroutine())

	// var wg sync.WaitGroup // 建立 WaitGroup 變數

	// // 單一 goroutine
	// go testWaitGroup(1, &wg)
	// wg.Add(1)

	// fmt.Println("啟動後 goroutines:", runtime.NumGoroutine())
	// wg.Wait() // 等待計數器歸零

	var wgwc sync.WaitGroup
	// 多個	goroutine
	ch1 := make(chan int)
	for i := 0; i < 10; i++ { // 建立 10 個 goroutine 針對同一個通道寫入
		wgwc.Add(1) // 增加計數器
		go testWaitGroupWithChannel(i, ch1, &wgwc)
	}

	go func() { // 因為需要等待所有 goroutine 完成，所以需要另外開一個 goroutine 來等待
		wgwc.Wait() // 等待計數器歸零
		close(ch1)
	}()

	fmt.Println("啟動後 goroutines:", runtime.NumGoroutine())
	for val := range ch1 { // 因為是通道阻塞，所以需要迴圈一個個讀取通道資料直到通道關閉
		fmt.Println("接收資料：", val)
	}

	fmt.Println("所有goroutine工作完成")
	fmt.Println("結束後 goroutines:", runtime.NumGoroutine())

	// fmt.Println("Hello, World!", externalAlias.ExportedValue, externalAlias.Myexternalfunc())
	// fmt.Println(quote.Go())
	// fmt.Println(globalVar)
	// fmt.Println("ccc", ccc)
	// f := scope()
	// v := f()
	// fmt.Println(v)

	// of, ov := outer()
	// fmt.Println(of(), ov)

	// type Vertex struct {
	// 	Lat, Long float64
	// 	T         int
	// }

	// vertexSlice := []Vertex{{40.68433, -74.39967, 333}}
	// fmt.Println(vertexSlice, cap(vertexSlice))

	// m := make(map[string]Vertex)

	// m["Bell Labs"] = Vertex{
	// 	40.68433, -74.39967, 333,
	// }
	// fmt.Println(m["Bell Labs"])

}

func scope() func() int {
	outer_var := 2
	foo := func() int { return outer_var }
	return foo
}

func outer() (func() int, int) {
	outer_var := 2
	inner := func() int {
		outer_var += 99 // outer_var from outer scope is mutated.
		return outer_var
	}
	inner()
	return inner, outer_var // return inner func and mutated outer_var 101
}
