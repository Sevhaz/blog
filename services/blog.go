package services

import (
	"blog/db"
	"blog/models"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BlogService struct{}

func (s *BlogService) ListAllBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	err := db.Db.Find(&blogs).Error
	return blogs, err
}

func (s *BlogService) GetBlog(blogID string) (*models.Blog, error) {
	var blog models.Blog
	err := db.Db.First(&blog, "id = ?", blogID).Error
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (s *BlogService) CreateBlog(blog *models.Blog, claims jwt.MapClaims) error {
	userID, ok := claims["userID"].(string)
	if !ok {
		return errors.New("invalid token claims")
	}

	blog.ID

