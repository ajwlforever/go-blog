package main

import (
	"fmt"
	"go-blog/dao"
	"testing"
)

func TestPostSelect(t *testing.T) {
	fmt.Println("sss")
	//initMysql()
	postDao := new(dao.PostDao)
	post := postDao.SelectById(1)
	fmt.Println(post.Title)
}
