package middleware

import (
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func AuthMiddleware(authService service.AuthService, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		tokenSlice := strings.Split(authHeader, " ")
		if len(tokenSlice) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		validateToken, err := authService.ValidateToken(tokenSlice[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		claims, ok := validateToken.Claims.(jwt.MapClaims)
		if !ok || !validateToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		userId := uint32(claims["user_id"].(float64))
		user, err := userService.FindById(context.Background(), userId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		c.Set("userInfo", user)
	}
}
