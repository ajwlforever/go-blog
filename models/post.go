package models

import (
	"fmt"
	"time"
)

type Post struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Desc         string    `json:"desc"`
	Flags        int       `json:"flags"`
	CreatedTime  time.Time `json:"created"`
	ModifiedTime time.Time `json:"modified"`
	IsDeleted    bool      `json:"is_deleted"`
	CreatorId    int64     `json:"creator_id"`
	Level        int       `json:"level"` // 置顶，等等等级设置

}

func (post *Post) ToString() string {
	return fmt.Sprintf("postId`%s`", post.Id)
}
