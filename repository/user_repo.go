package repository

import (
	"blog/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUserByID(userID string) (*models.User, error)
}

type UserRepo struct {
	Db *gorm.DB
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.Db.Where("Email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := r.Db.Create(user).Error
	return err
}

func (r *UserRepo) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := r.Db.Preload("Posts").Where("ID = ?", userID).First(&user).Error
	return &user, err
}
