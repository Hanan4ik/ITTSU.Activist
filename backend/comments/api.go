package comments

import (
	"activist/access"
	"activist/constants"
	"activist/cookie"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentsAPI struct {
	db     *CommentDB
	adb    *access.AcessDB
	cookie *cookie.CookieHandler
}

type StructPostID struct {
	PostID int64 `json:"postID"`
}

type StructID struct {
	ID int64 `json:"id"`
}

func NewCommentsAPI(db *CommentDB, adb *access.AcessDB, cookie *cookie.CookieHandler) CommentsAPI {
	return CommentsAPI{db, adb, cookie}
}

func (capi *CommentsAPI) IsStudent(g *gin.Context) int64 { // дописать json
	raw, err := g.Cookie("session")
	if err != nil {
		return 0
	}
	id, err := capi.cookie.ValidateCookie(raw)
	if err != nil {
		return 0
	}
	if capi.adb.IsStudent(id.ID) {
		return id.ID
	}
	return 0
}

func (capi *CommentsAPI) IsOrg(g *gin.Context) int64 { // дописать json
	raw, err := g.Cookie("session")
	if err != nil {
		return 0
	}
	id, err := capi.cookie.ValidateCookie(raw)
	if err != nil {
		return 0
	}
	if capi.adb.IsOrganisator(id.ID) {
		return id.ID
	}
	return 0
}

func (capi *CommentsAPI) HasRights(g *gin.Context, comment Comment) bool {
	raw, err := g.Cookie("session")
	if err != nil {
		return false
	}
	id, err := capi.cookie.ValidateCookie(raw)
	if err != nil {
		return false
	}
	if (comment.UserID == id.ID) || capi.adb.IsOrganisator(id.ID) {
		return true
	}
	return false
}

func (capi *CommentsAPI) AddComment(g *gin.Context) {
	var com Comment
	res := capi.IsStudent(g)
	if res == 0 {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return
	}
	g.BindJSON(&com)
	com.UserID = res
	capi.db.NewComment(com)
	g.JSON(http.StatusOK, constants.OKResponse)
}

func (capi *CommentsAPI) RemoveComment(g *gin.Context) {
	var id StructID
	g.BindJSON(&id)
	removeComment := capi.db.GetComment(id.ID)
	if !(capi.HasRights(g, removeComment)) {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return
	}
	capi.db.RemoveComment(id.ID)
	g.JSON(http.StatusOK, constants.OKResponse)
}

func (capi *CommentsAPI) GetComments(g *gin.Context) {
	var postID StructPostID
	g.BindJSON(&postID)
	removeComment := capi.db.GetPostComments(postID.PostID)
	g.JSON(http.StatusOK, removeComment)
}
