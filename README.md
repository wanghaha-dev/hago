# Hago Web framework

#### 介绍
hago 是一个零依赖，开箱即用的web框架，好了，就介绍这一点就完事了。
不明白的可以联系QQ: 31931727

#### 已支持的功能
中间件
分组路由
静态文件
HTML
...

#### 即将支持的功能
ORM
自动生成 增删改查 (CURD) 代码。
参数校验
JWT
...

#### 安装
```sh
go get github.com/wanghaha-dev/hago
```

#### 使用示例
```go
package main

import (
	"fmt"
	"github.com/wanghaha-dev/hago"
	"net/http"
)

func main() {
	server := hago.Default()
	server.GET("/hello", func(ctx *hago.Context) {
		ctx.String(http.StatusOK, "hello")
	})

	server.GET("/api/:name/list", func(ctx *hago.Context) {
		name := ctx.Param("name")
		age := ctx.Query("age")

		ctx.JSON(http.StatusOK, hago.H{
			"name": name,
			"age":  age,
		})
	})

	server.ANY("/any", func(ctx *hago.Context) {
		ctx.String(http.StatusOK, "any ...")
	})

	server.Static("/assets", "./static")

	server.LoadHTMLGlob("templates/*")

	server.GET("/h1", func(ctx *hago.Context) {
		ctx.HTML(http.StatusOK, "hello.html", nil)
	})

	v1 := server.Group("/v1")
	v1.GET("/user/list", func(ctx *hago.Context) {
		ctx.String(http.StatusOK, "v1 user list page...")
	})

	v1.Use(func(ctx *hago.Context) {
		fmt.Println("before...")
		ctx.Next()
		fmt.Println("after...")
	})

	server.Run(":9999")
}
```

#### 其他技术教学视频

https://space.bilibili.com/94634171?spm_id_from=333.1007.0.0