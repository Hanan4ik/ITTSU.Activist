// Package cookie Works with cookie
package cookie

import (
	"errors"
	"github.com/gorilla/securecookie"
	"time"
)

// Struct that works with cookies
type CookieHandler struct {
	Worker *securecookie.SecureCookie // securecookie struct to work with cookie
}

type Session struct {
	ID     int64 // id of user
	Expiry int64 // date of cookie expiry
}

// Generate new cookie with userID and expiry date of 1 week
func (cookie *CookieHandler) GenerateCookie(userID int64) (string, error) {
	s := Session{userID, time.Now().Add(7 * 24 * time.Hour).Unix()}
	encoded, err := cookie.Worker.Encode("session", s)
	return encoded, err
}

// Validate cookie and return userID if it's valid, return error otherwise
func (cookie *CookieHandler) ValidateCookie(inputCookie string) (Session, error) {
	var s Session
	err := cookie.Worker.Decode("session", inputCookie, &s)
	if err == nil {
		expiry := s.Expiry
		if expiry < time.Now().Unix() {
			return s, errors.New("Expired")
		}
		return s, nil
	}
	return s, err
}

// Create new CookieHandler
func NewHandler(hashKey []byte, blockKey []byte) CookieHandler {
	return CookieHandler{securecookie.New(hashKey, blockKey)}
}
