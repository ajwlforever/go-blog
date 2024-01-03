package dao

import (
	"fmt"
	"go-blog/models"
)

const (
	_updateFour = "update post  set title = ?, content = ? , intro = ?, flags = ? where id = ?"
	_selectAll  = "select * from post"
	_selectById = "select * from post where id = ?"
	_insert     = "insert into post(id,title,content,intro,flags,creator_id,level)" +
		" values (?,?,?,?,?,?,?)"
	_updateDeleted = "update post set is_deleted = ? where id = ?"
	_deleteFull    = "delete from post where id = ?"
	_countAll      = "select count(*) from post"
)

func (Dao *Dao) GetAllPostCount() (count int64) {
	err := db.Get(&count, _countAll)
	if err != nil {
		fmt.Printf("getAllPostCount failed:%v", err)
	}
	return
}
func (Dao *Dao) SelectAllPost() (posts []models.Post, err error) {
	posts = []models.Post{}
	err = db.Select(&posts, _selectAll)
	if err != nil {
		fmt.Printf("get all failed, err:%v\n", err)
		return
	}
	fmt.Printf("get all post successful~")
	return
}

func (Dao *Dao) SelectPostById(id int) (post models.Post) {

	err := db.Get(&post, _selectById, id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id : %d\n, postTitle: %s", id, post.Title)
	return
}

func (Dao *Dao) SavePost(post *models.Post) (res bool) {

	result, err := db.Exec(_insert, post.Id, post.Title, post.Content, post.Intro, post.Flags, post.CreatorId, post.Id)
	if err != nil {
		fmt.Printf("save failed, err:%v\n", err)
		return false
	}
	affectRows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("affectRows:%v \n", affectRows)
	return
}
func (Dao *Dao) UpdatePostById(post *models.Post) (affectedRow int64) {
	result, err := db.Exec(_updateFour, post.Title, post.Content, post.Intro, post.Flags, post.Id)
	if err != nil {
		fmt.Printf("UpdatePostById failed:%v", err)
	}
	affectedRow, err = result.RowsAffected()
	if err != nil {
		fmt.Printf("UpdatePostById failed, err:%v")
	}
	return
}
func (Dao *Dao) DeleteById(id int) (affectedRow int64) {
	result, err := db.Exec(_updateDeleted, 1, id) // 将is_deleted置1
	if err != nil {
		fmt.Printf("DeleteById failed%v", err)
	}
	affectedRow, err = result.RowsAffected()
	if err != nil {
		fmt.Printf("DeleteById failed, err:%v")
	}
	return

}
func (Dao *Dao) FullDeleteById(id int) (affectedRow int64) {
	result, err := db.Exec(_deleteFull, id)
	if err != nil {
		fmt.Printf("DeleteById failed%v", err)
	}
	affectedRow, err = result.RowsAffected()
	if err != nil {
		fmt.Printf("DeleteById failed, err:%v")
	}
	return
}
