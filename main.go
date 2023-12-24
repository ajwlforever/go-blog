package main

import (
	"fmt"
	"go-blog/config"
	"go-blog/router"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("start web")
	fmt.Printf("%v", config.Conf.App.Author)

	// sever服务开启
	server := http.Server{
		Addr: "127.0.0.1:8998",
	}
	// 地址拦截
	router.Router()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("server.ListenAndServe() Failed")
	}
}
