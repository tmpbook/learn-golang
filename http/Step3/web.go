package main

import (
	"io"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &MyHandler{},
		ReadTimeout: 6 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = hello
	mux["/bye"] = bye

	server.ListenAndServe()
}

type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
	}
	io.WriteString(w, "\nURL: "+r.URL.String())
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello module")
}

func bye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bye module")
}
