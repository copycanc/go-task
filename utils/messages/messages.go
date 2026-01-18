package messages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Сообщение "Неверные данные"
func IncorrectData(c *gin.Context, err error) {
	c.JSONP(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": "Неверные данные: " + err.Error(),
	})
}
