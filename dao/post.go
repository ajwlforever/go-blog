package dao

import (
	"fmt"
	"go-blog/models"
)

type PostDao struct {
}

func (postDao *PostDao) SelectAllPost() (posts []models.Post, err error) {
	sqlStr := "select * from post"
	err = db.Select(&posts, sqlStr)
	if err != nil {
		fmt.Printf("get all failed, err:%v\n", err)
		return
	}
	fmt.Printf("get all post successful~")
	return
}

func (postDao *PostDao) SelectById(id int) (post models.Post) {
	sqlStr := "select * from post where id = ?"
	err := db.Get(&post, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id : %d\n, postTitle: %s", id, post.Title)
	return
}

func (postDao *PostDao) SavePost(post models.Post) (res bool) {
	sqlStr := "insert into post(id,title,content,intro,flags,creator_id,level)" +
		" values (?,?,?,?,?,?,?)"
	result, err := db.Exec(sqlStr, post.Id, post.Title, post.Content, post.Intro, post.Flags, post.CreatorId, post.Id)
	if err != nil {
		fmt.Printf("save failed, err:%v\n", err)
		return false
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("save Id:%v \n", insertID)
	return true
}
