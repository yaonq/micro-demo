package main

import (
	"net/http"

	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
)

func main() {

	// 指定服务监听端口
	add := web.Address(":8001")

	// 创建web服务
	srv := web.NewService(add)

	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 192.168.100.10:8001?test=1
		log.Info(r.URL.Query().Get("test"), r.Form, r.PostFormValue("test"))

		w.Write([]byte("hello world"))
	})

	srv.Run()
}
