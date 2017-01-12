package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Password string `json:"-"`
}

func (User) TableName() string {
	return "users"
}

func FindUser(db *gorm.DB, id uint, password string) (*User, error) {
	var user User
	result := db.First(&user, id)

	if result.Error != nil {
		log.Panic(result.GetErrors())
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return &user, nil
	}

	return nil, nil
}
