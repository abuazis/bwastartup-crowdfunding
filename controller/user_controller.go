package controller

import (
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService: userService}
}

func (userController *userController) Register(c *gin.Context) {
	var request model.RegisterRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Meta: model.MetaResponse{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			},
			Data: nil,
		})
	}

	ctx := context.Background()
	response, err := userController.userService.Register(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Meta: model.MetaResponse{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Message: err.Error(),
			},
			Data: nil,
		})
	}

	// Success
	c.JSON(http.StatusOK, model.WebResponse{
		Meta: model.MetaResponse{
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
			Message: "Account has been registered",
		},
		Data: response,
	})
}
