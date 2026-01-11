package access

import (
	"activist/constants"

	"gorm.io/gorm"
)

type AcessDB struct {
	gormDB *gorm.DB
}

type Right struct {
	gorm.Model
	Rights uint `gorm:"not null"` // storing rights of userID
	UserID uint `gorm:"not null"` // userID
}

func NewAcessDB(db *gorm.DB) AcessDB {
	return AcessDB{gormDB: db}
}

func (adb *AcessDB) InitDB() error {
	return adb.gormDB.AutoMigrate(&Right{})
}

func (adb *AcessDB) AddRight(right Right) error {
	res := adb.gormDB.Create(&right)
	return res.Error
}

func (adb *AcessDB) IsStudent(userID uint) (bool, error) {
	rights, err := adb.GetUserRight(userID)
	if err == nil && rights.Rights >= constants.STUDENT {
		return true, nil
	}
	return false, err
}

func (adb *AcessDB) IsOrganisator(userID uint) (bool, error) {
	rights, err := adb.GetUserRight(userID)
	if rights.Rights >= constants.ORGANISATOR && err == nil {
		return true, nil
	}
	return false, err
}

func (adb *AcessDB) IsAdmin(userID uint) (bool, error) {
	rights, err := adb.GetUserRight(userID)
	if rights.Rights >= constants.ADMIN && err == nil {
		return true, nil
	}
	return false, err
}

func (adb *AcessDB) GetUserRight(userID uint) (Right, error) {
	var rights Right
	res := adb.gormDB.Where("user_id = ?", userID).Find(&rights)
	return rights, res.Error
}
