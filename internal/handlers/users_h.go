package handlers

import (
	"go-br-task/internal/models"
	"go-br-task/utils/messages"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Получить всеx пользователей
func (h *HandlerUser) GetUser(c *gin.Context) {
	user, httpStatus, err := h.userService.GetAllUser()
	if err != nil {
		messages.StatusHttpError(c, httpStatus, err)
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status": "OK",
		"users":  user,
	})
	return
}

// Создать нового пользователя
func (h *HandlerUser) CreateUser(c *gin.Context) {
	var user models.User
	if errr := c.ShouldBindJSON(&user); errr != nil {
		messages.StatusBadRequestDataH(c, errr)
		return
	}
	httpStatus, err := h.userService.EmailExist(user.Email)
	if httpStatus == 200 {
		if httpStatus, err = h.userService.CreateUser(user); err != nil {
			messages.StatusHttpError(c, httpStatus, err)
			return
		}
		messages.StatusHttpSuccess(c)
		return
	}
	messages.StatusHttpError(c, httpStatus, err)
	return
}

// Получить пользователя по ID
func (h *HandlerUser) GetUserID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatusE, errE := h.userService.UserExist(id)
	if httpStatusE == 200 {
		user, httpStatus, err := h.userService.GetUserID(id)
		if err != nil {
			messages.StatusHttpError(c, httpStatus, err)
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status": "ok",
			"task":   user,
		})
		return
	}
	messages.StatusHttpError(c, httpStatusE, errE)
	return
}

// Удалить пользователя по ID
func (h *HandlerUser) DeleteUserID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatus, err := h.userService.UserExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.userService.DeleteUserID(id)
		if err != nil {
			messages.StatusHttpError(c, httpStatus, err)
			return
		}
		messages.StatusHttpSuccess(c)
		return
	}
	messages.StatusHttpError(c, httpStatus, err)
	return
}

// Изменить пользователя
func (h *HandlerUser) UpdateUserID(c *gin.Context) {
	var chuser models.ChangeUser
	if errr := c.ShouldBindJSON(&chuser); errr != nil {
		messages.StatusBadRequestDataH(c, errr)
		return
	}
	id, _ := uuid.Parse(c.Param("id"))
	httpStatus, err := h.userService.UserExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.userService.UpdateUserID(id, chuser)
		if err != nil {
			messages.StatusHttpError(c, httpStatus, err)
			return
		}
		messages.StatusHttpSuccess(c)
		return
	}
	messages.StatusHttpError(c, httpStatus, err)
	return
}
