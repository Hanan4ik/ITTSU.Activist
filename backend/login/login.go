// Package login is responsible for managing login into personal account
package login

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginDB struct {
	Database *sql.DB
}

func (ldb *LoginDB) Init() { // username, email, password_salt, password_hash, 2fa bool,
	loginTableScheme := `
	CREATE TABLE creds(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(50) UNIQUE NOT NULL,
		tabel_number VARCHAR(16) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		password_salt VARHAR(16) NOT NULL,
		password_hash VARCHAR(256) NOT NULL
	);
	`
	ldb.Database.Exec(loginTableScheme)
}

func (ldb *LoginDB) AddCreds(username, email, password_salt, password_hash string) error {
	_, err := ldb.Database.Exec("INSERT INTO creds (username, email, password_salt, password_hash) VALUES (?,?,?,?)", username, email, password_salt, password_hash)
	return err
}

// VerifyUsername Verifies that password match username. Return {user_id, nil} if match and {-1, error} if otherwise
func (ldb *LoginDB) VerifyTableNumber(tabel_number string, password string) (int64, error) {

	res, err := ldb.Database.Query("SELECT id, password_salt, password_hash FROM creds WHERE tabel_number = ?", tabel_number)
	if err != nil {
		return -1, err
	}
	var passwordSalt, passwordHash string
	var id int64
	exist := res.Next()
	if !exist {
		return -1, errors.New("VerifyTableNumber: No such cred. res.Next() error")
	}
	err = res.Scan(&id, &passwordSalt, &passwordHash)
	if err != nil {
		return -1, err
	}
	saltedPassword := password + passwordSalt
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(saltedPassword))
	return id, err
}

func (ldb *LoginDB) VerifyEmail(email string, password string) (int64, error) {
	res, err := ldb.Database.Query("SELECT id, password_salt, password_hash FROM creds WHERE email = ?", email)
	if err != nil {
		return -1, err
	}
	var passwordSalt, passwordHash string
	var id int64
	exist := res.Next()
	if !exist {
		return -1, errors.New("VerifyEmail: No such cred. res.Next() error")
	}
	err = res.Scan(&id, &passwordSalt, &passwordHash)
	if err != nil {
		return -1, err
	}
	saltedPassword := password + passwordSalt
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(saltedPassword))
	return id, err
}
