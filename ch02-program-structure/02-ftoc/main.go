package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToc(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToc(boilingF))
}

// 第一个 float64 指的是 f 的类型，第二个是方法的类型，即返回值的类型
func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}

// "32°F = 0°C"
// "212°F = 100°C"
