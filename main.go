package main

import (
	"bwastartup-crowdfunding/controller"
	"bwastartup-crowdfunding/database"
	_ "bwastartup-crowdfunding/docs"
	"bwastartup-crowdfunding/middleware"
	"bwastartup-crowdfunding/repository"
	"bwastartup-crowdfunding/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Crowdfunding Web API
// @version 1.0
// @description Contains API for bwastartup-crowdfunding project

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	db := database.GetConnection()

	authService := service.NewAuthServiceImpl()

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userController := controller.NewUserController(userService, authService)

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", userController.Register)
		v1.POST("/sessions", userController.Login)
		v1.POST("/avatars", middleware.AuthMiddleware(authService, userService), userController.UploadAvatar)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatalln(r.Run(":8080"))
}
