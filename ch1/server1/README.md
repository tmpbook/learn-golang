

### 启动服务（加&为了后台运行）:
```
Ξ ch1/server1 git:(master) ▶ go run main.go &
[1] 63601
```

### 测试:
> 这里使用之前编写的 fetch 程序测试
```
Ξ ch1/server1 git:(master) ▶ go run ../fetch/main.go http://localhost:8000
URL.Path = "/"
Ξ ch1/server1 git:(master) ▶ go run ../fetch/main.go http://localhost:8000/hello
URL.Path = "/hello"
```