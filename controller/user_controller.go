package controller

import (
	"bwastartup-crowdfunding/entity"
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
	authService service.AuthService
}

func NewUserController(userService service.UserService, authService service.AuthService) *userController {
	return &userController{userService: userService, authService: authService}
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
func (campaignController *userController) Register(c *gin.Context) {
	var request model.RegisterRequest

	// Get request data
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
	_, err = campaignController.userService.CheckEmail(ctx, request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   "Email has been registered",
		})
		return
	}

	// Create user
	response, err := campaignController.userService.Register(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	// Generate JWT
	generateToken, err := campaignController.authService.GenerateToken(response.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}
	response.Token = generateToken

	// Success
	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

// Login godoc
// @Summary Login account
// @Description Login account use email and password.
// @ID login-user
// @Accept  json
// @Produce  json
// @Param LoginRequest body model.LoginRequest true "Login Account"
// @Success 200 {object} model.WebResponse{data=model.LoginResponse}
// @Failure 400 {object} model.WebResponse{data=string}
// @Failure 422 {object} model.WebResponse{data=[]string}
// @Failure 500 {object} model.WebResponse{data=string}
// @Router /sessions [post]
func (campaignController *userController) Login(c *gin.Context) {
	// Get request data
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

	// User login
	ctx := context.Background()
	response, err := campaignController.userService.Login(ctx, loginRequest)
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

	// Generate JWT
	generateToken, err := campaignController.authService.GenerateToken(response.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}
	response.Token = generateToken

	// Success
	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}

// UploadAvatar godoc
// @Summary Upload Avatar account
// @Description Upload image of avatar via form
// @ID upload-avatar
// @Accept  image/*
// @Produce  json
// @Param Authorization header string true "Token"
// @Success 200 {object} model.WebResponse
// @Failure 400 {object} model.WebResponse{data=string}
// @Failure 500 {object} model.WebResponse{data=string}
// @Router /avatars [post]
func (campaignController *userController) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	ctx := context.Background()
	userInfo := c.MustGet("userInfo").(entity.User)
	avatarFileName, err := campaignController.userService.SaveAvatar(ctx, userInfo.Id, file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	uploadDestination := "uploads/users/" + avatarFileName

	err = c.SaveUploadedFile(file, uploadDestination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	})
}
