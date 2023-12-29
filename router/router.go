package router

import (
	"go-blog/api"
	"net/http"
)

func Router() {
	// 地址拦截
	http.HandleFunc("/api/v1/getAllPost", api.API.GetAllPost)
	http.HandleFunc("/api/v1/savePost", api.API.SavePost)
}
