package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// 开始一个协程
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// 从channel ch 获取
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		// 传递到channel ch
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// 命令
// go run main.go https://www.baidu.com https://vuejs.org https://zhihu.com
// 结果
// 1.28s    8804 https://zhihu.com
// 2.77s     227 https://www.baidu.com
// 3.06s   11488 https://vuejs.org
// 3.06s elapsed
// tip: 哪个执行完哪个输出，总时间即最长运行
