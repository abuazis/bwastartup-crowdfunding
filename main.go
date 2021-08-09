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
	//gin.SetMode(gin.ReleaseMode)
	db := database.GetConnection()

	authService := service.NewAuthServiceImpl()

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userController := controller.NewUserController(userService, authService)

	campaignRepository := repository.NewCampaignRepositoryImpl(db)
	campaignService := service.NewCampaignServiceImpl(campaignRepository)
	campaignController := controller.NewCampaignController(campaignService)

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8mb max

	r.Static("/uploads/", "./uploads")

	v1 := r.Group("/api/v1")
	{
		// User
		v1.POST("/users", userController.Register)
		v1.POST("/sessions", userController.Login)
		v1.POST("/avatars", middleware.AuthMiddleware(authService, userService), userController.UploadAvatar)

		// Campaign
		v1.GET("/campaigns", campaignController.GetCampaigns)
		v1.GET("/campaigns/:id", campaignController.GetCampaignDetails)
		v1.POST("/campaigns", middleware.AuthMiddleware(authService, userService), campaignController.CreateCampaign)
		v1.PUT("/campaigns/:id", middleware.AuthMiddleware(authService, userService), campaignController.UpdateCampaign)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatalln(r.Run(":8080"))
}
