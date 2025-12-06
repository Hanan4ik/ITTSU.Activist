package register

import (
	"activist/access"
	"activist/constants"
	"activist/cookie"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChangeValueInput struct {
	ID int64
}

type InputCreds struct {
	TabelNumber string
	Username    string
	Email       string
	Password    string
	Rights      int
}

type RegAPI struct {
	rdb    *RegisterDB
	adb    *access.AcessDB
	cookie *cookie.CookieHandler
}

func NewAPI(rdb *RegisterDB, adb *access.AcessDB, cookie *cookie.CookieHandler) RegAPI {
	return RegAPI{rdb, adb, cookie}
}

func (rAPI *RegAPI) GetInfo(g *gin.Context) { // /api/register/getInfo
	rawCookie, err := g.Cookie("session")
	if err != nil {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return
	}
	res, err := rAPI.cookie.ValidateCookie(rawCookie)
	if err != nil {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return
	}
	if !rAPI.adb.IsAdmin(res.ID) {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return
	}
	rows, err := rAPI.rdb.GetInfo()
	if err != nil {
		g.JSON(http.StatusInternalServerError, constants.ServerError)
		return
	}
	g.JSON(http.StatusOK, rows)
}

func (rAPI *RegAPI) Approve(g *gin.Context) { // /api/register/approve
	var input ChangeValueInput
	g.Bind(&input)
	id := input.ID
	err := rAPI.rdb.Approve(id)
	if err != nil {
		fmt.Println(err.Error())
		g.JSON(http.StatusInternalServerError, constants.ServerError)
	}
	g.JSON(http.StatusOK, constants.OKResponse)
}

func (rAPI *RegAPI) Disapprove(g *gin.Context) { // /api/register/disapprove
	var input ChangeValueInput
	g.Bind(&input)
	id := input.ID
	rAPI.rdb.Disapprove(id)
	g.JSON(http.StatusOK, constants.OKResponse)
}

func (rAPI *RegAPI) Register(g *gin.Context) { // api/register/create
	var input InputCreds
	g.Bind(&input)
	err := rAPI.rdb.QueueAppend(input.Username, input.TabelNumber, input.Email, input.Password, input.Rights)
	if err != nil {
		fmt.Println(err.Error())
	}
}
