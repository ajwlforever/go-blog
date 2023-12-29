package api

import (
	"fmt"
	"net/http"
)

func (api *Api) GetAllPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 鉴权

}

func (api *Api) SavePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := r.GetBody()
	if err != nil {
		fmt.Println("err:", err)
	}
	var params []byte
	str, err := body.Read(params)
	if err != nil {
		fmt.Println("err:", err)
	}
	println(str)
}
