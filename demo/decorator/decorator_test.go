package decorator

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func Hello(s string) {
	fmt.Println(s)
}
func Test(t *testing.T) {

	decorator(Hello)("Hello, World!")

}

func Test_compute(t *testing.T) {

	sum1 := timedSumFunc(Sum1)
	sum2 := timedSumFunc(Sum2)
	fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved Request %s from %s\n", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}

func Test_http_server(t *testing.T) {
	http.HandleFunc("/hello", WithServerHeader(hello))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Test_http_server_enhance(t *testing.T) {
	http.HandleFunc("/v4/hello", Handler(hello,
		WithAuthCookie, WithServerHeader, WithBasicAuth, WithDebugLog))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
