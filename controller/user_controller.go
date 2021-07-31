package controller

import (
	"bwastartup-crowdfunding/exception"
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/service"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
		c.JSON(http.StatusUnprocessableEntity, model.WebResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: http.StatusText(http.StatusUnprocessableEntity),
			Data:   exception.ValidationError(err),
		})
		return
	}

	ctx := context.Background()
	response, err := userController.userService.Register(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   nil,
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

func (userController *userController) Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, model.WebResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: http.StatusText(http.StatusUnprocessableEntity),
			Data:   exception.ValidationError(err),
		})
		return
	}

	ctx := context.Background()
	response, err := userController.userService.Login(ctx, loginRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Data:   "Wrong email",
			})
		} else if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			c.JSON(http.StatusUnauthorized, model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Data:   "Wrong password",
			})
		} else {
			c.JSON(http.StatusInternalServerError, model.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: http.StatusText(http.StatusInternalServerError),
				Data:   nil,
			})
		}
		return
	}

	// Success
	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}
