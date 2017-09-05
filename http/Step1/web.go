package main

import (
	"io"
	"net/http"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "hello tmpbook")
}

func main() {
	// HandleFunc 源代码如下
	http.HandleFunc("/", hello)
	// 此处省略了错误处理，如果出错会自动退出，比如端口占用
	http.ListenAndServe(":8080", nil)
}

// func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
// 	DefaultServeMux.HandleFunc(pattern, handler)
// }
