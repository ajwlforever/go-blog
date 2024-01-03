package api

import (
	"encoding/json"
	"fmt"
	"go-blog/services"
	"net/http"
)

func (api *Api) GetAllPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	service := services.New()

	posts, err := service.GetAllPosts()
	if err != nil {
		fmt.Printf("service.GetAllPosts() Failed")
	}
	resStr, err := json.Marshal(posts)
	if err != nil {
		fmt.Printf("json.Marshal(indexData) Failed")
	}
	w.Write(resStr)
}
