package login

import (
	"activist/cookie"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginAPI struct {
	Ldb    *LoginDB
	Cookie *cookie.CookieHandler
}

type LoginInput struct {
	Cred     string `json:"login"`
	Password string `json:"password"`
}

func isEmail(cred string) bool {
	return strings.Contains(cred, "@")
}

func (lapi LoginAPI) Login(c *gin.Context) { // /api/login endpoint
	var input LoginInput
	c.BindJSON(&input)
	fmt.Println(input)
	cred := input.Cred
	password := input.Password
	var id int64
	var err error
	fmt.Println(cred)
	if isEmail(cred) {
		id, err = lapi.Ldb.VerifyEmail(cred, password)
	} else {
		id, err = lapi.Ldb.VerifyTableNumber(cred, password)
	}
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"HUETA": "EBANNAYA"}) // Переделать
		fmt.Print(err.Error())
		return
	}
	cook, err := lapi.Cookie.GenerateCookie(id)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"HUETA": "EBANNAYA"}) // Переделать
		fmt.Print(err.Error())
		return
	}
	c.Header("Set-Cookie", fmt.Sprintf("session=%s; Path=/; Domain=localhost; Max-Age=86400; Secure; HttpOnly; SameSite=None; Partitioned", cook)) //
	c.JSON(http.StatusOK, gin.H{"OK": "OK"})
}

func (lapi LoginAPI) CheckCookie(c *gin.Context) {
	fmt.Println(c.Cookie("session"))
}
