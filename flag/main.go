package main

import (
	"flag"
	"fmt"
)

var (
	// Return a point which can read by *name after flag.Parse() executed
	name = flag.String("name", "unknow", "your name")
)

func main() {
	flag.Parse()

	// Create a visitor function, execute with a *flag.Flag callback param for each flag
	visitor := func(f *flag.Flag) {
		fmt.Println("option =", f.Name, " value =", f.Value)
	}
	flag.VisitAll(visitor)

	fmt.Println(*name)
}

// cmd   : go run main.go -name tmpbook
// output:
// 	   option = name  value = tmpbook
// 	   tmpbook

// cmd   : go run main.go
// output:
// 	   option = name  value = unknow
// 	   unknow
