package constants

import "github.com/gin-gonic/gin"

var (
	STUDENT     uint = 0
	ORGANISATOR uint = 1
	ADMIN       uint = 2
)

var (
	AuthErr     = gin.H{"error": "Not authorized", "ok": ""}
	OKResponse  = gin.H{"error": "", "ok": "Yes"}
	ServerError = gin.H{"error": "Internal service error", "ok": ""}
	WrongCreds  = gin.H{"error": "Wrong username or password", "ok": "ok"}
)
