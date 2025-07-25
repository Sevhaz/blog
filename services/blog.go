package services

import (

	"blog/models"
	"blog/repository"
	

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BlogService struct {
	Repo repository.BlogRepository
}

func (s *BlogService) ListAllPosts(claims jwt.MapClaims) ([]models.Blog, error) {
	userId, err := uuid.Parse(claims["userID"].(string))
	if err != nil {
		return nil ,  err
	}
	posts, err := s.Repo.GetAllPosts(userId)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *BlogService) GetPost(id uuid.UUID) (*models.Blog, error) {
	blog, err := s.Repo.GetBlogByID(id)
	if err != nil {
		return &models.Blog{}, err
	}
	return blog, nil
}

func (s *BlogService) CreateBlogPost(req *models.Blog, claims jwt.MapClaims) error {
	userId, err := uuid.Parse(claims["userID"].(string))
	if err != nil {
		return err
	}
	req.AuthorID = userId
	err = s.Repo.CreateBlog(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogService) UpdateBlog(req *models.Blog, blogID string) error {
	id:= uuid.MustParse(blogID)
	blog, err := s.Repo.GetBlogByID(id)
	if err != nil {
		return err
	}
	
	if req.Title != ""{
	blog.Title = req.Title	
	}
	if req.Content != ""{
	blog.Content = req.Content
	}
	
	err = s.Repo.SaveBlog(blog)
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogService) DeleteBlogPost(blogID string) error {
	id:= uuid.MustParse(blogID)

	err := s.Repo.DeleteBlogPost(id)
	if err != nil {
		return err
	}
	return nil
}