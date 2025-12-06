package events

import (
	"activist/access"
	"activist/constants"
	"activist/cookie"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventAPI struct {
	db     *EventDB
	cookie *cookie.CookieHandler
	adb    *access.AcessDB
}

func NewEventAPI(db *EventDB, cookie *cookie.CookieHandler, adb *access.AcessDB) EventAPI {
	return EventAPI{db, cookie, adb}
}

type Event struct {
	ID       int64  `json:"id"` // -1 if creating
	Title    string `json:"title"`
	Text     string `json:"text"`
	Location string `json:"location"`
}

func (eapi *EventAPI) IsOrg(g *gin.Context) bool {
	rawCookie, err := g.Cookie("session")
	fmt.Println(rawCookie)
	if err != nil {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return false
	}
	res, err := eapi.cookie.ValidateCookie(rawCookie)
	if err != nil {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return false
	}
	id := res.ID
	if !eapi.adb.IsOrganisator(id) {
		g.JSON(http.StatusBadRequest, constants.AuthErr)
		return false
	}
	return true
}

func (eapi *EventAPI) CreateEvent(g *gin.Context) {
	var event Event
	g.BindJSON(&event)
	if !eapi.IsOrg(g) {
		return
	}
	eapi.db.CreateEvent(event.Title, event.Text, event.Location)
	g.JSON(http.StatusOK, constants.OKResponse)
}

func (eapi *EventAPI) RemoveEvent(g *gin.Context) {
	var id int64
	g.BindJSON(&id)
	if !eapi.IsOrg(g) {
		return
	}
	eapi.db.RemoveEvent(id)
	g.JSON(http.StatusOK, constants.OKResponse)
}

func (eapi *EventAPI) GetEvents(g *gin.Context) {
	scheme := "SELECT * FROM events"
	res, err := eapi.db.GetEvents(scheme)
	if err != nil {
		fmt.Println(err.Error())
		g.JSON(http.StatusInternalServerError, constants.ServerError)
		return
	}
	g.JSON(http.StatusOK, res)
}

func (eapi *EventAPI) UpdateEvent(g *gin.Context) {
	var event Event
	g.BindJSON(&event)
	if !eapi.IsOrg(g) {
		return
	}
	eapi.db.UpdateEvent(event)
	g.JSON(http.StatusOK, constants.OKResponse)
}
