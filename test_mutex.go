package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func muxtexTest(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	count++
	fmt.Println(count)
	mu.Unlock()
}
