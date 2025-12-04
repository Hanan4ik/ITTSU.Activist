package register

import (
	crypto_back "activist/crypto"
	"database/sql"
)

type RegisterDB struct {
	db *sql.DB
}

type Row struct {
	ID           int64
	Username     string
	Table_number string
	Email        string
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
	approved INT DEFAULT 0
	);
	`
	rdb.db.Exec(scheme)
}

func (rdb *RegisterDB) QueueAppend(username, table_number, email, password string) error {
	salt_bytes, err := crypto_back.GenerateSalt(16)
	if err != nil {
		return err
	}
	password_salt := crypto_back.BytesToBase64(salt_bytes)
	salted_password := password + password_salt
	password_hash, err := crypto_back.HashPassword(salted_password)
	if err != nil {
		return err
	}
	_, err = rdb.db.Exec("INSERT INTO register username, table_number, email, password_salt, password_hash", username, table_number, email, password_salt, password_hash)
	return err

}

func (rdb *RegisterDB) GetInfo() ([]Row, error) {
	infoRes := make([]Row, 0)
	res, err := rdb.db.Query("SELECT username, table_number, email FROM register WHERE approved = 0")
	if err != nil {
		return infoRes, err
	}
	for res.Next() {
		var newRow Row
		res.Scan()
		err = res.Scan(&newRow.ID, &newRow.Username, &newRow.Table_number, &newRow.Email)
		if err != nil {
			return infoRes, err
		}
		infoRes = append(infoRes, newRow)
	}
	return infoRes, nil
}
