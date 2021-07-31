package main

import (
	"bwastartup-crowdfunding/controller"
	"bwastartup-crowdfunding/database"
	"bwastartup-crowdfunding/repository"
	"bwastartup-crowdfunding/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db := database.GetConnection()
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userController := controller.NewUserController(userService)

	r := gin.Default()
	api := r.Group("/api/v1")

	api.POST("/users", userController.Register)
	api.POST("/sessions", userController.Login)
	log.Fatalln(r.Run())
}
