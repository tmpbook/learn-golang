### 代码执行过程

1. 调用 `Http.HandleFunc`
    1. 调用 `DefaultServeMux` 的 `HandleFunc`
    2. 调用 `DefaultServeMux` 的 `Handle`
    3. 向 `DefaultServeMux` 的 `map[string]muxEntry` 中添加对应的 handler 和路由规则
2. 调用 `http.ListenAndServe(":8080", nil)`
    1. 实例化 Server
    2. 调用 Server 的 `ListenAndServe()`
    3. 调用 `net.Listen("tcp", addr)` 监听端口
    4. 启动一个 for 循环，在循环体中 Accept 请求
    5. 对每个请求实例化一个 Conn，并且开启一个 goroutine 为这个请求进行服务： `go c.serve()`
    6. 读取每个请求的内容 `w, err := c.readRequest()`
    7. 判断 handler 是否为空，如果没有设置 handler，handler就设置为 `DefaultServeMux`
    8. 调用 handler 的 ServeHTTP
    9. 根据 request 选择 handler，并且进入到这个 handler 的 ServeHTTP

            mux.handler(r).ServeHTTP(w, r)
    
    10. 选择 handler：
        1. 判断是否有路由能满足这个 request（循环遍历 ServeMux 的 muxEntry）
        2. 如果有路由满足，调用路由 handler 的ServeHTTP
        3. 如果没有路由满足，调用 NotFoundHandler 的 ServeHTTP

    > 小时候想当科学家是觉得科学家可以改变世界，现在我发现，敲键盘也可以