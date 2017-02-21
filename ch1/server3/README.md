### Request:
```
http://localhost:8000/?name=kevin&age=18
```
### Response:
```
GET /?name=kevin&age=18 HTTP/1.1
Header["Cache-Control"] = ["max-age=0"]
Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"]
Header["Accept-Encoding"] = ["gzip, deflate, sdch, br"]
Header["Accept-Language"] = ["zh-CN,zh;q=0.8,en;q=0.6"]
Header["Cookie"] = ["connect.sid=s%3AK-oRJfbgI15fOg8i3YZ-moUOdVGoIOTa.bmM%2FRTp19zKMZCKWQOtYWBKP11%2FUJhiCtPYfrPZeBys"]
Header["Connection"] = ["keep-alive"]
Header["Upgrade-Insecure-Requests"] = ["1"]
Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"]
Host = "localhost:8000"
RemoteAddr = "127.0.0.1:61706"
Form["name"] = ["kevin"]
Form["age"] = ["18"]
```