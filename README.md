# Cross-Origin Resource Sharing [![GoDoc](https://godoc.org/github.com/go-mego/cors?status.svg)](https://godoc.org/github.com/go-mego/cors)

Cross-Origin Resource Sharing 為跨來源資源共用，這讓你能夠在不同網站向你的伺服器發送請求。通常來說這不建議使用，因為請求應該要侷限於相同來源以避免惡意攻擊或偽造，但在開發時則有可能需要此功能。

# 索引

* [安裝方式](#安裝方式)
* [使用方式](#使用方式)
	* [預設配置](#預設配置)

# 安裝方式

打開終端機並且透過 `go get` 安裝此套件即可。

```bash
$ go get github.com/go-mego/cors
```

# 使用方式

將 `cors.New` 傳入 Mego 引擎的 `Use` 來將跨來源資源共用中介軟體用於所有路由裡。

```go
package main

import (
	"github.com/go-mego/cors"
	"github.com/go-mego/mego"
)

func main() {
	m := mego.New()
	// 套用跨來源資源共用中介軟體到每個路由。
	m.Use(cors.New(&cors.Options{
		AllowOrigins: []string{"https://*.example.com"},
		AllowMethods: []string{"PUT", "PATCH"},
		AllowHeaders: []string{"Origin"},
	}))
	m.Run()
}
```

跨來源資源共用中介軟體也能夠僅用於單個路由上，以避免整個伺服端都呈現公開可跨域狀態。

```go
func main() {
	m := mego.New()
	// 跨來源資源共用中介軟體也能夠僅套用到單一路由，
	// 這令你能夠更彈性地替不同路由配置不同設定。
	m.GET("/", cors.New(&cors.Options{
		// ...
	}), func() {
		// ...
	})
	m.Run()
}
```

## 預設配置

透過 `Default` 可以使用預設的跨域配置，這會允許所有來源都能進行跨域連線。這是最不安全的設置但同時也很適合用於開發的寬鬆環境，在正式且公開的伺服器上請絕對不要使用此設置。

```go
func main() {
	m := mego.New()
	// `Default` 會使用最不安全的跨域設置，但這在開發環境十分方便。
	m.Use(cors.Default())
	m.Run()
}
```