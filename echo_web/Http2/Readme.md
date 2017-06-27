生成`HTTP/2`需要的`cert.pem`和`key.pem`文件

```
go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
```