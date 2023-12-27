package main

import (
	"fmt"
	"go-blog/dao"
	"go-blog/models"
	"testing"
)

func TestPostSelect(t *testing.T) {
	fmt.Println("sss")
	//initMysql()
	postDao := new(dao.PostDao)
	post := postDao.SelectById(1)
	fmt.Println(post.Title)
}

func TestSavePost(t *testing.T) {
	postDao := new(dao.PostDao)
	var post models.Post
	post.Title = "123"
	post.Id = 12
	post.Content = "CCCCCC"
	post.Intro = "1111"
	post.Flags = "falgs"
	post.CreatorId = 1
	post.Level = 1

	res := postDao.SavePost(post)

	fmt.Printf("res: %v", res)
}

func TestSelectAllPost(t *testing.T) {
	postDao := new(dao.PostDao)
	posts, err := postDao.SelectAllPost()

	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(posts)
}
