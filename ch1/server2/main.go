package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 用来输出请求链接的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}

// 计算并输出到现在为止 call 的次数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// 这里访问"/count"的时候并不会触发"/"的handler
// 这个服务器有两个请求处理函数,根据请求的 url 不
// 同会调用不同的函数:对/count 这个 url 的请求会
// 调用到 counter 这个函数,其它的 url 都会调用默
// 认的处理函数。如果你的请求 pattern 是以/结尾,那
// 么所有以该 url 为前缀的 url 都会被这条规则匹配。
// 在这些代码的背后, 服务器每一次接收请求处理时都会
// 另起一个 goroutine,这样服务器就可以同一时间处理
// 多个 请求。然而在并发情况下,假如真的有两个请求同
// 一时刻去更新 count,那么这个值可能并不 会被正确地
// 增加;这个程序可能会引发一个严重的 bug:竞态条件
// (参见 9.1)。为了避免这 个问题,我们必须保证每次
// 修改变量的最多只能有一个 goroutine,这也就是代码
// 里的 mu.Lock()和 mu.Unlock()调用将修改 count
// 的所有行为包在中间的目的。
