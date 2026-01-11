package comments

import (
	"activist/access"
	"activist/constants"
	"activist/cookie"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentsAPI struct {
	db     *CommentDB
	adb    *access.AcessDB
	cookie *cookie.CookieHandler
}

type StructPostID struct {
	PostID uint `json:"postID"`
}

type StructID struct {
	ID uint `json:"id"`
}

func NewCommentsAPI(db *CommentDB, adb *access.AcessDB, cookie *cookie.CookieHandler) CommentsAPI {
	return CommentsAPI{db, adb, cookie}
}

func (capi *CommentsAPI) IsStudent(g *gin.Context) (uint, error) { // дописать json
	raw, err := g.Cookie("session")
	if err != nil {
		return 0, err
	}
	id, err := capi.cookie.ValidateCookie(raw)
	if err != nil {
		return 0, err
	}
	res, err := capi.adb.IsStudent(id.ID)
	if err == nil && res {
		return id.ID, nil
	}
	return 0, err
}

func (capi *CommentsAPI) IsOrg(g *gin.Context) (uint, error) { // дописать json
	raw, err := g.Cookie("session")
	if err != nil {
		return 0, err
	}
	id, err := capi.cookie.ValidateCookie(raw)
	if err != nil {
		return 0, err
	}
	res, err := capi.adb.IsOrganisator(id.ID)
	if err == nil && res {
		return id.ID, nil
	}
	return 0, errors.New("NO RESULT")
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
	res, err := capi.adb.IsOrganisator(id.ID)
	if ((comment.UserID == id.ID) || res) && err == nil {
		return true
	}
	return false
}

func (capi *CommentsAPI) AddComment(g *gin.Context) {
	var com Comment
	res, err := capi.IsStudent(g)
	if res == 0 {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return
	}
	if err != nil {
		g.JSON(http.StatusInternalServerError, constants.ServerError)
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
	removeComment, err := capi.db.GetComment(id.ID)
	if !(capi.HasRights(g, removeComment)) || err != nil {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return
	}
	capi.db.RemoveComment(id.ID)
	g.JSON(http.StatusOK, constants.OKResponse)
}

func (capi *CommentsAPI) GetComments(g *gin.Context) {
	var postID StructPostID
	g.BindJSON(&postID)
	removeComment, err := capi.db.GetPostComments(postID.PostID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, constants.ServerError)
	}
	g.JSON(http.StatusOK, removeComment)
}
