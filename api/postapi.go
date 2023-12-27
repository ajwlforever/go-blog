package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *Api) GetAllPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	indexData := IndexData{Title: "sss", Desc: "desc"}

	resStr, err := json.Marshal(indexData)
	if err != nil {
		fmt.Printf("json.Marshal(indexData) Failed")
	}
	w.Write(resStr)
}
