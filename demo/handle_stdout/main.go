package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// func runSomeShell(args ...string) {
func runSomeShell(args []string) {
	cmd := exec.Command("sh", args...)
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
	// runSomeShell("./echo.sh", "1", "2", "3")
	runSomeShell([]string{"./echo.sh", "1", "2", "3"})
}
