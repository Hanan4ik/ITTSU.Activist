package comments

import (
	"os"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB(t *testing.T) CommentDB {
	sqlDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("CANNOT OPEN DB")
	}
	db := NewCommentDB(sqlDB)
	err = db.Init()
	if err != nil {
		t.Fatalf("CANNOT INIT DB")
	}
	return db
}

func TestInit(t *testing.T) {
	InitDB(t)
	defer os.Remove("test.db")
}

func TestCRUD(t *testing.T) {
	db := InitDB(t)
	defer os.Remove("test.db")
	// CREATE
	com := Comment{Text: "I love lusik", PostID: 1, UserID: 1}
	for range 10 {
		err := db.NewComment(com)
		if err != nil {
			t.Error("ERROR CREATING NEW COMMENT")
		}
	}

	// READ
	comments, err := db.GetPostComments(1)
	for i := range 10 {
		if comments[i].Text != com.Text || err != nil {
			t.Error("ERROR READING POSTS COMMENT")
		}
	}
	t.Log(comments)

	nthComment := comments[3]
	idComment, err := db.GetComment(nthComment.ID)
	if idComment != nthComment || err != nil {
		t.Error("ERROR READING COMMENT BY ID")
	}

	// UPDATE
	for i := range 10 {
		comments[i].Text = "My code is awesome"
		err := db.UpdateComment(comments[i].ID, comments[i])
		if err != nil {
			t.Error("ERROR UPDATING COMMENT")
		}
	}

	newComments, err := db.GetPostComments(1)
	for i := range 10 {
		if newComments[i].Text != comments[i].Text {
			t.Error("ERROR IN UPDATED COMMENT")
		}
	}

	// DELETE
	for i := range 10 {
		err := db.RemoveComment(newComments[i].ID)
		if err != nil {
			t.Error("ERROR IN REMOVING COMMENT")
		}
	}
	removedComments, _ := db.GetPostComments(1)
	if len(removedComments) != 0 {
		t.Error("ERROR, COMMENTS WEREN'T REMOVED")
	}

}
