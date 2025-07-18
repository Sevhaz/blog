package routes

import (
	"blog/handlers"
	"blog/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, blogHandler *handlers.BlogHandler) *mux.Router {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Protected routes with middleware
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/user/info", userHandler.LoginInfo).Methods("GET")

	protected.HandleFunc("/blogs", blogHandler.ListAllBlogs).Methods("GET")
	protected.HandleFunc("/blogs/{blogID}", blogHandler.BlogByID).Methods("GET")
	protected.HandleFunc("/blogs", blogHandler.CreateBlog).Methods("POST")
	protected.HandleFunc("/blogs/{blogID}", blogHandler.UpdateBlog).Methods("PUT")
	protected.HandleFunc("/blogs/{blogID}", blogHandler.DeleteBlog).Methods("DELETE")

	return r
}
