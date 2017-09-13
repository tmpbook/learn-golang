package main

import (
	"fmt"
)

// func new(Type) *Type

// 和 build-in 方法：new(int) 一样
func newInt() *int {
	var i int
	return &i
}
func newString() *string {
	var i string
	return &i
}

func main() {
	myNewInt := newInt()
	fmt.Println(*myNewInt)

	myNewString := newString()
	fmt.Println(*myNewString)
}
