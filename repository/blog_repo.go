package repository

import (
	"blog/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) Create(blog *models.Blog) error {
	return r.db.Create(blog).Error
}

func (r *BlogRepository) GetByID(id uuid.UUID) (*models.Blog, error) {
	var blog models.Blog
	err := r.db.First(&blog, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *BlogRepository) ListByUserID(userID uuid.UUID) ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Where("user_id = ?", userID).Find(&blogs).Error
	return blogs, err
}

func (r *BlogRepository) Update(blog *models.Blog) error {
	return r.db.Save(blog).Error
}

func (r *BlogReposi*
