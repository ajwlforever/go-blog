package main

import (
	"fmt"
	"go-blog/dao"
	"go-blog/models"
	"testing"
	"time"
)

func TestPostSelect(t *testing.T) {
	fmt.Println("sss")
	//initMysql()
	postDao := new(dao.PostDao)
	post := postDao.SelectPostById(1)
	fmt.Println(post.Title)
}

func TestSavePost(t *testing.T) {
	postDao := new(dao.PostDao)
	var post models.Post
	post.Title = "123"
	post.Id = 1233
	post.Content = "CCCCCC"
	post.Intro = "1111"
	post.Flags = "falgs"
	post.CreatorId = 1
	post.Level = 1
	post.CreatedTime = time.Now()
	post.ModifiedTime = time.Now()
	res := postDao.SavePost(&post)

	fmt.Printf("res: %v", res)
}

func TestSelectAllPost(t *testing.T) {
	postDao := new(dao.PostDao)
	posts, err := postDao.SelectAllPost()

	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(len(posts))
	fmt.Println(posts)
}

func TestAll1(t *testing.T) {
	postDao := new(dao.PostDao)
	fmt.Println(postDao.GetAllPostCount())
	postDao.FullDeleteById(123)
	fmt.Println(postDao.GetAllPostCount())
	var post models.Post
	post.Title = "123"
	post.Id = 123
	post.Content = "CCCCCC"
	post.Intro = "1111"
	post.Flags = "falgs"
	post.CreatorId = 1
	post.CreatedTime = time.Now()
	post.ModifiedTime = time.Now()
	post.Level = 1
	postDao.SavePost(&post)
	fmt.Println(postDao.GetAllPostCount())
	postDao.DeleteById(123)
	fmt.Printf("is_deleted:%v", postDao.SelectPostById(123))
	post.Title = "llll"
	post.Id = 12
	post.Content = "cccc"
	post.Intro = "aaaa"
	post.Flags = "bbbbb"
	postDao.UpdatePostById(&post)

}
