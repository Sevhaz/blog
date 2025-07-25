package handlers

import (
	"blog/models"
	"blog/services"
	"blog/utils"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
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
		http.Error(w, "user not found", http.StatusInternalServerError)
		return
	}
	

	blogPosts, err := h.Service.ListAllPosts(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blogPosts)
}
func (h *BlogHandler) UpdateBlogPost(w http.ResponseWriter, r *http.Request) {
	var req models.Blog
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params:= mux.Vars(r)
	blogid := params["id"]
	err = h.Service.UpdateBlog(&req, blogid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("blog successfully updated")
}


func (h *BlogHandler) DeleteBlogPost(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
	id:= vars["id"]
	
	err := h.Service.DeleteBlogPost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("blog successfully deleted")
}

