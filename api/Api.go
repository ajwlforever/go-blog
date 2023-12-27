package api

var API = &Api{}

type Api struct {
}

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
