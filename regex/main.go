package main

import (
	"fmt"
	"regexp"
)

func checkUserName(myname string) (rt bool) {
	if m, _ := regexp.MatchString(`^([a-zA-Z]+)\.([a-zA-Z]+)$`, myname); !m {
		rt = false
	} else {
		rt = true
	}
	return
}

func main() {
	fmt.Println(checkUserName("KevinGao"))
	fmt.Println(checkUserName("Kevin.Gao"))
}
