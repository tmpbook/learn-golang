package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func runSomeShell(name string, args ...string) {
	// func runSomeShell(name string, args []string) {
	cmd := exec.Command(name, args...)
	var out, err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err
	error := cmd.Run()
	if error != nil {
		fmt.Println("Run failed:\n", err.String())
		return
	}
	fmt.Printf(out.String())
}

func main() {
	runSomeShell("ansible", "jumpserver", "-m ping", "-u root")
	runSomeShell("sh", "./echo.sh", "a", "b", "c")
	// runSomeShell("sh", []string{"./echo.sh", "a", "b", "c"})
}

// go run main.go
// ----
// 107.*.*.33 | SUCCESS => {
//     "changed": false,
//     "ping": "pong"
// }
// 第一个参数为: a 参数个数为: 3
// 第一个参数为: b 参数个数为: 2
// 第一个参数为: c 参数个数为: 1
