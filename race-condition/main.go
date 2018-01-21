package main

import (
	"fmt"
	"sync"
)

// N 是内存里的一个数子
var (
	N         = 0
	waitgroup sync.WaitGroup
)

func counter(number *int) {
	*number++
	waitgroup.Done()
}

func main() {

	for i := 0; i < 1000; i++ {
		waitgroup.Add(1)
		go counter(&N)
	}
	waitgroup.Wait()
	fmt.Println(N)
}
