package main

import (
	"fmt"
	"sync"
)

// N 是内存里的一个对象
var (
	N         = 0
	waitgroup sync.WaitGroup
)

func counter(number *int) {
	for i := 0; i < 1000; i++ {
		*number++
	}
	waitgroup.Done()
}

func main() {
	waitgroup.Add(2)
	go counter(&N)
	go counter(&N)

	waitgroup.Wait()
	fmt.Println(N)
}
