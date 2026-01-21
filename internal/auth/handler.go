package auth

import (
	"go-br-task/internal/message"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type HandlerAuth struct {
	authService *AuthService
	jwtSecret   []byte
}

func NewHandlerAuth(authService *AuthService, jwtSecret string) *HandlerAuth {
	return &HandlerAuth{authService: authService, jwtSecret: []byte(jwtSecret)}
}

func (h *HandlerAuth) Login(c *gin.Context) {
	var newAuth ForAuth
	if err := c.ShouldBindJSON(&newAuth); err != nil {
		message.StatusBadRequestDataH(c, err)
		return
	}
	id, httpStatus, err := h.authService.CheckCredentials(newAuth.Email, newAuth.Password)
	if err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	claims := jwt.MapClaims{
		"sub":  id.String(),
		"name": newAuth.Email,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(h.jwtSecret)
	if err != nil {
		message.StatusHttpError(c, 500, err)
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"status": "OK",
		"token":  tokenStr,
	})

}
