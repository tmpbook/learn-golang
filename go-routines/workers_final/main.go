package main

import "fmt"

var N int = 100
var R int = 100

// Task 任务
func Task(i int) {
	fmt.Println("Box", i)
}

// Workers worker
func Workers(task func(interface{}), climax func()) chan interface{} {
	input := make(chan interface{})
	ack := make(chan bool)
	for i := 0; i < R; i++ {
		go func() {
			for {
				v, ok := <-input
				if ok {
					task(v)
					ack <- true
				} else {
					return
				}
			}
		}()
	}
	go func() {
		for i := 0; i < R; i++ {
			<-ack
		}
		climax()
	}()
	return input
}

func main() {

	exit := make(chan bool)

	workers := Workers(func(a interface{}) {
		Task(a.(int))
	}, func() {
		exit <- true
	})

	for i := 0; i < N; i++ {
		workers <- i
	}
	close(workers)

	<-exit
}
