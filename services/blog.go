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

// func (s *BlogService) GetPost(id string) (*models.Blog, error) {
// 	blog, err := s.Repo.GetPostByID(id)
// 	if err != nil {
// 		return &models.Blog{}, err
// 	}
// 	return blog, nil
// }

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

// func (s *BlogService) UpdatePost(req *models.Blog, blogID string, claims jwt.MapClaims) error {
// 	blog, err := s.Repo.GetPostByID(blogID)
// 	if err != nil {
// 		return err
// 	}
// 	userId, err := uuid.Parse(claims["userID"].(string))
// 	if err != nil {
// 		return err
// 	}
// 	if blog.AuthorID != userId {
// 		return errors.New("unable edit others posts")
// 	}

// 	blog.Title = req.Title
// 	blog.Content = req.Content
// 	req = blog
// 	err = s.Repo.SaveBlog(req)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *BlogService) DeletePost(claims jwt.MapClaims, id string) error {
// 	post, err := s.Repo.GetPostByID(id)
// 	if err != nil {
// 		return err
// 	}
// 	userId, err := uuid.Parse(claims["userID"].(string))
// 	if err != nil {
// 		return err
// 	}
// 	if post.AuthorID != userId {
// 		return errors.New("unable to delete blogs")
// 	}

// 	err = database.Db.Delete(post).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }