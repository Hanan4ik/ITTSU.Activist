package comments

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text   string `gorm:"not null"`
	PostID int64  `gorm:"not null"`
	UserID int64  `gorn:"not null"`
}

type CommentDB struct {
	gorm *gorm.DB
}

func NewCommentDB(dbPath string) (CommentDB, error) {
	var res CommentDB
	gdb, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return res, err
	}
	res = CommentDB{gdb}
	return res, nil
}

func (cmd *CommentDB) Init() {
	cmd.gorm.AutoMigrate(&Comment{})
}

func (cmd *CommentDB) NewComment(comment Comment) {
	cmd.gorm.Create(&comment)
}

func (cmd *CommentDB) GetComment(id int64) Comment {
	var name Comment
	cmd.gorm.Where("id = ?", id).Find(&Comment{}).Scan(&name)
	return name
}

func (cmd *CommentDB) GetPostComments(id int64) []Comment {
	var comments []Comment
	cmd.gorm.Where("post_id = ?", id).Find(&comments)
	fmt.Println(comments)
	return comments
}

func (cmd *CommentDB) UpdateComment(id int64, comment Comment) {
	cmd.gorm.Where("id = ?", id).Updates(&comment)
}

func (cmd *CommentDB) RemoveComment(id int64) {
	fmt.Println(id)
	cmd.gorm.Delete(&Comment{}, id)
}
