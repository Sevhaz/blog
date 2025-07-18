package main

import (
	"blog/database"
	"blog/handlers"
	"blog/repository"
	"blog/routes"
	"blog/services"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	/* initialise database */
	database.OpenDb()

	/* initialise repositories */
	userRepo := &repository.UserRepo{}
	blogRepo := &repository.BlogRepo{}

	/* initialise services */
	userService := &services.UserService{Repo: userRepo}
	blogService := &services.BlogService{Repo: blogRepo}

	/* initialise handlers */
	userHandler := &handlers.UserHandler{Service: userService}
	blogHandler := &handlers.BlogHandler{Service: blogService}

	/* initialise routes */
	r := routes.SetUpRouter(userHandler, blogHandler)

	/* start server */
	fmt.Println("starting server...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}
}
