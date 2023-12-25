package dao

import (
	"fmt"
	"go-blog/models"
)

type PostDao struct {
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
