package main

import "fmt"

var N int = 100

func Task(i int) {
	fmt.Println("Box", i)
}
func main() {
	ack := make(chan bool, N) // Acknowledgement channel
	for i := 0; i < N; i++ {
		go func(arg int) { // Point #1
			Task(arg)
			ack <- true // Point #2
		}(i) // Point #3
	}

	for i := 0; i < N; i++ {
		<-ack // Point #2
	}
}
