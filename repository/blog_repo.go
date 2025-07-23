package repository

import (
	"blog/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogRepository interface {
	GetAllPosts(userID uuid.UUID) ([]models.Blog, error)
	GetBlogByID(id string) (*models.Blog, error)
	CreateBlog(post *models.Blog) error
	SaveBlog(post *models.Blog) error
}

type BlogRepo struct {
	Db *gorm.DB
}

func (r *BlogRepo) GetAllPosts(userID uuid.UUID) ([]models.Blog, error) {
	var posts []models.Blog
	err := r.Db.Where("author_id = ?", userID).Find(&posts).Error
	return posts, err

}

func (r *BlogRepo) GetBlogByID(id string) (*models.Blog, error) {
	var post models.Blog
	err := r.Db.Preload("Author").Where("ID = ?", id).First(&post).Error
	return &post, err
}

func (r *BlogRepo) CreateBlog(post *models.Blog) error {
	err := r.Db.Create(post).Error
	return err
}

func (r *BlogRepo) SaveBlog(post *models.Blog) error {
	err := r.Db.Save(post).Error
	return err
}
