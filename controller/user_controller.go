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

// Register godoc
// @Summary Register account
// @Description Create account with name, occupation, email, and password data
// @ID register-user
// @Accept  json
// @Produce  json
// @Param RegisterRequest body model.RegisterRequest true "Register Account"
// @Success 200 {object} model.WebResponse{data=model.RegisterResponse}
// @Failure 400 {object} model.WebResponse{data=string}
// @Failure 422 {object} model.WebResponse{data=[]string}
// @Router /users [post]
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

	// Check Email
	ctx := context.Background()
	_, err = userController.userService.CheckEmail(ctx, request.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   "Email has been registered",
		})
		return
	}
	response, err := userController.userService.Register(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
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
				Data:   err.Error(),
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
