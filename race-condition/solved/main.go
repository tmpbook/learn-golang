package main

import (
	"fmt"
	"sync"
)

// N 是内存里的一个对象
var (
	N         = 0
	mutex     sync.Mutex // +
	waitgroup sync.WaitGroup
)

func counter(number *int) {
	mutex.Lock() // ++
	*number++
	mutex.Unlock() // +++
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
