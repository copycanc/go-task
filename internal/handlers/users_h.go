package handlers

import (
	"go-br-task/internal/models"
	"go-br-task/utils/messages"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Получить всеx пользователей
func (h *HandlerUser) GetUser(c *gin.Context) {
	user, err := h.userService.GetAllUser()
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"status": "ok",
		"users":  user,
	})
	return
}

// Создать нового пользователя
func (h *HandlerUser) CreateUser(c *gin.Context) {
	var user models.User
	if errr := c.ShouldBindJSON(&user); errr != nil {
		messages.IncorrectData(c, errr)
		return
	}
	httpStatus, err := h.userService.EmailExist(user.Email)
	if httpStatus == 200 {
		if err = h.userService.CreateUser(user); err != nil {
			c.JSONP(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		c.JSONP(http.StatusOK, gin.H{
			"status": "ok",
			"text":   "Пользователь создан",
		})
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
	return
}

// Получить пользователя по ID
func (h *HandlerUser) GetUserID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatusE, errE := h.userService.UserExist(id)
	if httpStatusE == 200 {
		user, httpStatus, err := h.userService.GetUserID(id)
		if err != nil {
			c.JSONP(httpStatus, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status": "ok",
			"task":   user,
		})
		return
	}
	c.JSONP(httpStatusE, gin.H{
		"status":  "error",
		"message": errE.Error(),
	})
	return
}

// Удалить пользователя по ID
func (h *HandlerUser) DeleteUserID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatus, err := h.userService.UserExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.userService.DeleteUserID(id)
		if err != nil {
			c.JSONP(httpStatus, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status":  "ok",
			"message": "Пользователь удален",
		})
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
	return
}

// Изменить пользователя
func (h *HandlerUser) UpdateUserID(c *gin.Context) {
	var chuser models.ChangeUser
	if errr := c.ShouldBindJSON(&chuser); errr != nil {
		messages.IncorrectData(c, errr)
		return
	}
	id, _ := uuid.Parse(c.Param("id"))
	httpStatus, err := h.userService.UserExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.userService.UpdateUserID(id, chuser)
		if err != nil {
			c.JSONP(httpStatus, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status":  "ok",
			"message": "Данные обновлены",
		})
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
	return
}
