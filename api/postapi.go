package api

import (
	"encoding/json"
	"fmt"
	"go-blog/models"
	"go-blog/services"
	"io"
	"log"
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

func (api *Api) SavePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println("service.SavePost: %v", r.Method)
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Unable to read")
			// 处理读取请求体错误
			http.Error(w, "无法读取请求体", http.StatusInternalServerError)
			return
		}
		// 处理请求体
		// fmt.Println("请求体内容:", string(body))
		post := &models.Post{}
		err = json.Unmarshal(body, post)
		if err != nil {
			log.Println("请求体解码失败:", err)
		}
		fmt.Println("POST:", post)
		defer r.Body.Close()
		w.Write(Result(1, "SUCCESS"))
		return

	}
	w.Write(Result(404, "请求方式错误"))
	return
}
