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
	Dao := new(dao.Dao)
	post := Dao.SelectPostById(1)
	fmt.Println(post.Title)
}
func Test1(t *testing.T) {
	var post *models.Post

	post.Id = 1
	fmt.Printf("%v", post)
}

func TestSavePost(t *testing.T) {
	Dao := new(dao.Dao)

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
	res := Dao.SavePost(&post)

	fmt.Printf("res: %v", res)
}

func TestSelectAllPost(t *testing.T) {
	Dao := new(dao.Dao)
	posts, err := Dao.SelectAllPost()

	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(len(posts))
	fmt.Println(posts)
}

func TestAll1(t *testing.T) {
	Dao := new(dao.Dao)
	fmt.Println(Dao.GetAllPostCount())
	Dao.FullDeleteById(123)
	fmt.Println(Dao.GetAllPostCount())
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
	Dao.SavePost(&post)
	fmt.Println(Dao.GetAllPostCount())
	Dao.DeleteById(123)
	fmt.Printf("is_deleted:%v", Dao.SelectPostById(123))
	post.Title = "llll"
	post.Id = 12
	post.Content = "cccc"
	post.Intro = "aaaa"
	post.Flags = "bbbbb"
	Dao.UpdatePostById(&post)

}
