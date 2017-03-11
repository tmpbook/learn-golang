package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var seq = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *seq))
	if !*n {
		fmt.Println()
	}
}

// 使用 flag.Bool 函数会创建一个新的对应布尔型标志参数的变量。它有三个属性：第一个是命令行标志参数的名字『n』
// 然后是该标志参数的默认值『false』,最后是该标志参数对应的描述信息。如果用户在命令行输入了一个无效的标志参数，
// 或者输入 -h，-help参数，那么将打印出来，测试用例如下：
//
// $ go build gopl.io/ch2/echo4
// $ ./echo4 a bc def a bc def
// $ ./echo4 -s / a bc def a/bc/def
// $ ./echo4 -n a bc def a bc def$
// $ ./echo4 -help
// Usage of ./echo4:
//   -n    omit trailing newline
//   -s string
//         separator (default " ")
