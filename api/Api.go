package api

import (
	"encoding/json"
	"fmt"
)

var API = &Api{}

type Api struct {
}

type Res struct {
	Code int
	Msg  string
}

func Result(code int, msg interface{}) (result []byte) {
	resStr, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("err:%v", err)
	}
	var res = Res{
		Code: code,
		Msg:  string(resStr),
	}
	result, err = json.Marshal(res)
	if err != nil {
		fmt.Println("err:%v", err)
	}
	return
}
