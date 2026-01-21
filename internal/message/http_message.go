package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusBadRequestDataH(c *gin.Context, err error) {
	c.JSONP(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": "неверные данные " + err.Error(),
	})
}

func StatusHttpError(c *gin.Context, httpStatus int, err error) {
	c.JSONP(httpStatus, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
}

func StatusHttpSuccess(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"status": "OK",
	})
}
