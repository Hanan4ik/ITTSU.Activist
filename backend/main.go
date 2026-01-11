package main

import (
	"activist/access"
	"activist/comments"
	"activist/cookie"
	crypto_back "activist/crypto"
	"activist/events"
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

func startAPI(lapi login.LoginAPI, rapi *register.RegAPI, eapi *events.EventAPI, capi *comments.CommentsAPI) {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"} // Add your frontend origin
	config.AllowCredentials = true                          // Important for cookies
	router.Use(cors.New(config))
	router.GET("/api/register/getInfo", rapi.GetInfo)
	router.POST("/api/register/approve", rapi.Approve)
	router.POST("/api/register/disapprove", rapi.Disapprove)
	router.POST("/api/register/create", rapi.Register)
	router.POST("/api/login", lapi.Login)
	router.POST("/api/check", lapi.CheckCookie)
	router.POST("/api/events/getEvents", eapi.GetEvents)
	router.POST("/api/events/createEvent", eapi.CreateEvent)
	router.POST("/api/events/removeEvent", eapi.RemoveEvent)
	router.POST("/api/events/updateEvent", eapi.UpdateEvent)
	router.POST("/api/comments/get", capi.GetComments)
	router.POST("/api/comments/create", capi.AddComment)
	router.POST("/api/comments/remove", capi.RemoveComment)
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
	loginDB.AddCreds("org", "9999", "org@org", password_salt, password_hash)
	eventDB := events.NewDB(db)
	eventDB.Init()
	eventAPI := events.NewEventAPI(&eventDB, &cookie, &accesDB)
	accesDB.AddRight(1, 2)
	accesDB.AddRight(2, 1)
	err = regDB.QueueAppend("pidorok", "2", "b@b", "password", 0)
	cdb, _ := comments.NewCommentDB("comments.sqlite")
	cdb.Init()
	cdb.NewComment(comments.Comment{Text: "Privet mir", PostID: 3, UserID: 1})
	cdb.NewComment(comments.Comment{Text: "Privet mir", PostID: 3, UserID: 1})
	capi := comments.NewCommentsAPI(&cdb, &accesDB, &cookie)
	fmt.Println(cdb.GetComment(1))
	if err != nil {
		fmt.Print(err.Error())
	}
	startAPI(loginAPI, &registerAPI, &eventAPI, &capi)
}
