package main

import (
	"io"
	"net/http"
)

type MyHandle struct{}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &MyHandle{})
	http.ListenAndServe(":8080", mux)
}

func (*MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: "+r.URL.String())
}
