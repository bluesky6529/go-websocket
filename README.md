# go-websocket
### 參考文檔
```
https://ithelp.ithome.com.tw/articles/10237376
```

#### for jenkins using so not support color module

### remember windows to linux env
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build .
```

## how to use
```
go run main.go ws://abc.com wss://abc.com
```