package router

import (
	"go-blog/api"
	"net/http"
)

func Router() {
	// 地址拦截
	http.HandleFunc("/", api.API.Index)
}
