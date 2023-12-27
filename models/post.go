package models

import (
	"fmt"
	"time"
)

type Post struct {
	Id           int64     `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Content      string    `json:"content" db:"content"`
	Intro        string    `json:"intro" db:"intro"`
	Flags        string    `json:"flags" db:"flags"`
	CreatedTime  time.Time `json:"created_time" db:"created_time"`
	ModifiedTime time.Time `json:"modified_time" db:"modified_time"`
	IsDeleted    bool      `json:"is_deleted" db:"is_deleted"`
	CreatorId    int64     `json:"creator_id" db:"creator_id"`
	Level        int       `json:"level" db:"level"` // 置顶，等等等级设置

}

func (post *Post) ToString() string {
	return fmt.Sprintf("postId`%s`", post.Id)
}
