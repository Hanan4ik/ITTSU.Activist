// Package login is responsible for managing login into personal account
package login

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginDB struct {
	Database *gorm.DB
}

type LoginRecord struct {
	gorm.Model
	Username     string `gorm:"not null, unique"`
	TabelNumber  uint   `gorm:"not null, unique"`
	Email        string `gorm:"not null, unique"`
	PasswordSalt string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
}

func (ldb *LoginDB) Init() error {
	err := ldb.Database.AutoMigrate(&LoginRecord{})
	return err
}

func (ldb *LoginDB) AddCreds(record *LoginRecord) (uint, error) {
	result := ldb.Database.Create(&record)
	err := result.Error
	if err != nil {
		return 0, err
	}
	return uint(record.ID), err
}

// VerifyUsername Verifies that password match username. Return {user_id, nil} if match and {0, error} otherwise
func (ldb *LoginDB) VerifyTableNumber(tabel_number string, password string) (uint, error) {
	var result LoginRecord
	res := ldb.Database.Where("tabel_number = ?", tabel_number).First(&result)
	err := res.Error
	if err != nil {
		return 0, err
	}
	passwordSalt := result.PasswordSalt
	passwordHash := result.PasswordHash
	saltedPassword := password + passwordSalt
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(saltedPassword))
	return result.ID, err
}

func (ldb *LoginDB) VerifyEmail(email string, password string) (uint, error) {
	var result LoginRecord
	res := ldb.Database.Where("email = ?", email).First(&result)
	err := res.Error
	if err != nil {
		return 0, err
	}
	passwordSalt := result.PasswordSalt
	passwordHash := result.PasswordHash
	saltedPassword := password + passwordSalt
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(saltedPassword))
	return result.ID, err
}
