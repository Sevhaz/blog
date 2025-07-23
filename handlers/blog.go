package handlers

import (
	"blog/models"
	"blog/services"
	"blog/utils"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	_"github.com/gorilla/mux"
)

type BlogHandler struct {
	Service *services.BlogService
}


func (h *BlogHandler) CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	var req models.Blog
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}

	err = h.Service.CreateBlogPost(&req, claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("blog successfully created")
}

func (h *BlogHandler) GetBlogs(w http.ResponseWriter, r *http.Request) {
	
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}

	

	blogPosts, err := h.Service.ListAllPosts(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blogPosts)
}

// func (h *BlogHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
// 	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
// 	if !ok {
// 		http.Error(w, "user not found in context", http.StatusInternalServerError)
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	userID := vars["userID"]

// 	err := h.Service.DeleteBlog(claims, userID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
