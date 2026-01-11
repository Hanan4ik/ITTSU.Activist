package comments

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text   string `gorm:"not null"`
	PostID uint   `gorm:"not null"`
	UserID uint   `gorn:"not null"`
}

type CommentDB struct {
	gorm *gorm.DB
}

func NewCommentDB(gdm *gorm.DB) CommentDB {
	return CommentDB{gorm: gdm}
}

func (cmd *CommentDB) Init() error {
	return cmd.gorm.AutoMigrate(&Comment{})
}

func (cmd *CommentDB) NewComment(comment Comment) error {
	res := cmd.gorm.Create(&comment)
	return res.Error
}

func (cmd *CommentDB) GetComment(id uint) (Comment, error) {
	var name Comment
	res := cmd.gorm.Where("id = ?", id).Find(&Comment{}).Scan(&name)
	return name, res.Error
}

func (cmd *CommentDB) GetPostComments(id uint) ([]Comment, error) {
	var comments []Comment
	res := cmd.gorm.Where("post_id = ?", id).Find(&comments)
	return comments, res.Error
}

func (cmd *CommentDB) UpdateComment(id uint, comment Comment) error {
	res := cmd.gorm.Where("id = ?", id).Updates(&comment)
	return res.Error
}

func (cmd *CommentDB) RemoveComment(id uint) error {
	res := cmd.gorm.Delete(&Comment{}, id)
	return res.Error
}
