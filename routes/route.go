package routes

import (
	"blog/handlers"
	"blog/middleware"

	"github.com/gorilla/mux"
)

func SetUpRouter(userHandler *handlers.UserHandler, postHandler *handlers.BlogHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.Register).Methods("POST") // '/register' this is a route that leads you to the appropiate handler to handle
	// the incoming request, this is the same for all routes made
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	p := r.PathPrefix("/").Subrouter()
	p.Use(middleware.AuthMiddleware)

	//p.HandleFunc("/me", userHandler.LoginInfo).Methods("GET")

	p.HandleFunc("/posts", postHandler.GetBlogs).Methods("GET")
	//p.HandleFunc("/posts/{id}", postHandler.BlogByID).Methods("GET")
	p.HandleFunc("/posts", postHandler.CreateBlogPost).Methods("POST")
	//p.HandleFunc("/posts/{id}",postHandler.UpdateBlog).Methods("PUT")
	//p.HandleFunc("/posts/{id}",postHandler.DeleteBlog).Methods("DELETE")

	return r
}
