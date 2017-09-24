package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("boring!")
	d := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("you say: %q\n", <-c)
		fmt.Printf("you say: %q\n", <-d)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
