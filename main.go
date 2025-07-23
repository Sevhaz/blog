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
	db := database.OpenDb() //initializing your database and opening a connection so your application can interact with your database

	/* initialise repositories */
	userRepo := &repository.UserRepo{Db: db} //contains all user interactions with database e.g create user, update get user etc
	blogRepo := &repository.BlogRepo{Db: db} //contains all blog interactions with database e.g. create blog, update blog and get blog etc

	/* initialise service */
	userService := &services.UserService{Repo: userRepo} //this handles the applications and business logic required fot users
	blogService := &services.BlogService{Repo: blogRepo} //this handles the applications and business logic required fot blog

	/* initialise handlers */
	userHandler := &handlers.UserHandler{Service: userService} //handles all user requests eg login, create user etc
	postHandler := &handlers.BlogHandler{Service: blogService} // handles all blog posts eg. create new post, get post etc

	/* initialise routes */
	r := routes.SetUpRouter(userHandler, postHandler) //handles all incoming requests and directs them to the appriate route

	/* start server */
	fmt.Println("starting server...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r) //port is entry point into the server, where the application is listening - network ports are between 0 to 65,535
	if err != nil {                                      //listen and serve starts the application on a speicfic port which is taken from your env file
		log.Fatal("failed to start server:", err) //if there is a error it will print the error and stop the application
	}
}
