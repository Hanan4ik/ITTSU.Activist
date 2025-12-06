package register

import (
	"activist/access"
	crypto_back "activist/crypto"
	"activist/login"
	"database/sql"
	"fmt"
)

type RegisterDB struct {
	db  *sql.DB
	ldb *login.LoginDB
	adb *access.AcessDB
}

type Row struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	TableNumber string `json:"tableNumber"`
	Email       string `json:"email"`
	Rights      int64  `json:"rights"`
}

type InternalRow struct {
	Row          Row
	PasswordSalt string
	PasswordHash string
}

func (rdb *RegisterDB) Init() {
	scheme := `
	CREATE TABLE register(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(50) UNIQUE NOT NULL,
	table_number VARCHAR(16) UNIQUE NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	password_salt VACHAR(16) NOT NULL,
	password_hash VACHAR(256) NOT NULL,
	rights INT NOT NULL, 
	approved INT DEFAULT 0
	);
	`
	rdb.db.Exec(scheme)
}

func NewRDB(db *sql.DB, ldb *login.LoginDB, adb *access.AcessDB) RegisterDB {
	return RegisterDB{db, ldb, adb}
}

func (rdb *RegisterDB) QueueAppend(username, tableNumber, email, password string, rights int) error {
	saltBytes, err := crypto_back.GenerateSalt(16)
	if err != nil {
		return err
	}
	passwordSalt := crypto_back.BytesToBase64(saltBytes)
	saltedPassword := password + passwordSalt
	passwordHash, err := crypto_back.HashPassword(saltedPassword)
	if err != nil {
		return err
	}
	_, err = rdb.db.Exec("INSERT INTO register (username, table_number, email, password_salt, password_hash, rights, approved) VALUES (?,?,?,?,?,?,?)", username, tableNumber, email, passwordSalt, passwordHash, rights, 0)
	return err
}

func (rdb *RegisterDB) GetInfo() ([]Row, error) {
	req := "SELECT id, username, table_number, email, rights FROM register WHERE approved = 0;"
	res, err := rdb.QueryInfo(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	return res, err
}

func (rdb *RegisterDB) Approve(id int64) error {
	req := "SELECT id, username, table_number, email, rights, password_salt, password_hash FROM register WHERE id = ?"
	res, err := rdb.QueryInternalInfo(req, id)
	if err != nil {
		return err
	}
	row := res[0]
	userID, err := rdb.ldb.AddCreds(row.Row.Username, row.Row.TableNumber, row.Row.Email, row.PasswordSalt, row.PasswordHash)
	if err != nil {
		return err
	}
	rdb.adb.AddRight(userID, row.Row.Rights)
	rdb.ChangeApproveValue(id, 1)
	return nil
}

func (rdb *RegisterDB) Disapprove(id int64) {
	rdb.ChangeApproveValue(id, -1)
}

func (rdb *RegisterDB) ChangeApproveValue(id, value int64) {
	rdb.db.Exec("UPDATE register SET approved = ? WHERE id = ?", value, id)
}

func (rdb *RegisterDB) QueryInfo(scheme string, args ...any) ([]Row, error) {
	infoRes := make([]Row, 0)
	res, err := rdb.db.Query(scheme, args...)
	if err != nil {
		return infoRes, err
	}
	for res.Next() {
		var newRow Row
		err = res.Scan(&newRow.ID, &newRow.Username, &newRow.TableNumber, &newRow.Email, &newRow.Rights)
		if err != nil {
			return infoRes, err
		}
		infoRes = append(infoRes, newRow)
	}
	return infoRes, nil
}

func (rdb *RegisterDB) QueryInternalInfo(scheme string, args ...any) ([]InternalRow, error) {
	infoRes := make([]InternalRow, 0)
	res, err := rdb.db.Query(scheme, args...)
	if err != nil {
		return infoRes, err
	}
	for res.Next() {
		var newRow InternalRow
		err = res.Scan(&newRow.Row.ID, &newRow.Row.Username, &newRow.Row.TableNumber, &newRow.Row.Email, &newRow.Row.Rights, &newRow.PasswordSalt, &newRow.PasswordHash)
		if err != nil {
			return infoRes, err
		}
		infoRes = append(infoRes, newRow)
	}
	return infoRes, nil
}
