package services

import (
	"errors"
	"library/db"
	"library/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) Register(user *models.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	return db.Db.Create(user).Error
}

func (s *UserService) Login(user *models.User) (string, error) {
	var foundUser models.User
	err := db.Db.Where("email = ?", user.
