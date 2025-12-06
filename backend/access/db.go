package access

import (
	"activist/constants"
	"database/sql"
)

type AcessDB struct {
	db *sql.DB
}

func NewAcessDB(db *sql.DB) AcessDB {
	return AcessDB{db}
}

func (adb *AcessDB) InitDB() {
	scheme := `
	CREATE TABLE rights(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	rights INTEGER NOT NULL,
	FOREIGN KEY (id) REFERENCES creds(id)
	);`
	adb.db.Exec(scheme)
}

func (adb *AcessDB) AddRight(userID, rights int64) {
	adb.db.Exec("INSERT INTO rights VALUES(?,?)", userID, rights)
}

func (adb *AcessDB) IsStudent(userID int64) bool {
	rights, err := adb.GetUserRight(userID)
	if rights == constants.STUDENT && err == nil {
		return true
	}
	return false
}

func (adb *AcessDB) IsOrganisator(userID int64) bool {
	rights, err := adb.GetUserRight(userID)
	if rights == constants.ORGANISATOR && err == nil {
		return true
	}
	return false
}

func (adb *AcessDB) IsAdmin(userID int64) bool {
	rights, err := adb.GetUserRight(userID)
	if rights == constants.ADMIN && err == nil {
		return true
	}
	return false
}

func (adb *AcessDB) GetUserRight(userID int64) (int64, error) {
	var rights int64
	res, err := adb.db.Query("SELECT rights FROM rights WHERE id = ?", userID)
	if err != nil {
		return -1, err
	}
	res.Next()
	res.Scan(&rights)
	return rights, nil
}
