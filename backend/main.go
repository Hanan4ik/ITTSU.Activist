package main

import (
	"activist/access"
	"activist/cookie"
	crypto_back "activist/crypto"
	"activist/login"
	"activist/register"
	"crypto/rand"
	"database/sql"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/turso-go"
)

func startAPI(lapi login.LoginAPI, rapi *register.RegAPI) {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:3000"} // Add your frontend origin
	config.AllowCredentials = true                          // Important for cookies
	router.Use(cors.New(config))
	router.POST("/api/register/getInfo", rapi.GetInfo)
	router.POST("/api/register/approve", rapi.Approve)
	router.POST("/api/register/disapprove", rapi.Disapprove)
	router.POST("/api/register/create", rapi.Register)
	router.POST("/api/login", lapi.Login)
	router.POST("/api/check", lapi.CheckCookie)
	router.Run("localhost:8000")
}

func randomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func main() {
	db, err := sql.Open("turso", "db.sqlite")
	regDb, err := sql.Open("turso", "rdb.sqlite")
	if err != nil {
		fmt.Print(err.Error())
	}
	hashKey := randomBytes(64)
	blockKey := randomBytes(32)
	cookie := cookie.NewHandler(hashKey, blockKey)
	accesDB := access.NewAcessDB(db)
	accesDB.InitDB()
	loginDB := login.LoginDB{db}
	loginDB.Init()
	loginAPI := login.LoginAPI{&loginDB, &cookie}
	regDB := register.NewRDB(regDb, &loginDB, &accesDB)
	regDB.Init()
	registerAPI := register.NewAPI(&regDB, &accesDB, &cookie)
	password_salt := "1234"
	password := "password" + password_salt
	password_hash, _ := crypto_back.HashPassword(password)
	loginDB.AddCreds("admin", "0", "a@a", password_salt, password_hash)
	accesDB.AddRight(1, 2)
	err = regDB.QueueAppend("pidorok", "2", "b@b", "password", 0)
	if err != nil {
		fmt.Print(err.Error())
	}
	startAPI(loginAPI, &registerAPI)
}
