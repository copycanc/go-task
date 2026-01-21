package user

import (
	"go-br-task/internal/message"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HandlerUser struct {
	userService *UserService
}

func NewHandlerUser(userService *UserService) *HandlerUser {
	return &HandlerUser{userService: userService}
}

// Получить всеx пользователей
func (h *HandlerUser) List(c *gin.Context) {
	user, httpStatus, err := h.userService.GetAllUser()
	if err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status": "OK",
		"users":  user,
	})
	return
}

// Создать нового пользователя
func (h HandlerUser) Create(c *gin.Context) {
	var user User
	if errr := c.ShouldBindJSON(&user); errr != nil {
		message.StatusBadRequestDataH(c, errr)
		return
	}
	httpStatus, err := h.userService.EmailExist(user.Email)
	if httpStatus == 200 {
		if httpStatus, err = h.userService.CreateUser(user); err != nil {
			message.StatusHttpError(c, httpStatus, err)
			return
		}
		message.StatusHttpSuccess(c)
		return
	}
	message.StatusHttpError(c, httpStatus, err)
	return
}

// Получить пользователя по ID
func (h *HandlerUser) Get(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatusE, errE := h.userService.UserExist(id)
	if httpStatusE == 200 {
		user, httpStatus, err := h.userService.GetUserID(id)
		if err != nil {
			message.StatusHttpError(c, httpStatus, err)
			return
		}
		c.JSONP(httpStatus, gin.H{
			"status": "ok",
			"task":   user,
		})
		return
	}
	message.StatusHttpError(c, httpStatusE, errE)
	return
}

// Удалить пользователя по ID
func (h *HandlerUser) Delete(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	httpStatus, err := h.userService.UserExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.userService.DeleteUserID(id)
		if err != nil {
			message.StatusHttpError(c, httpStatus, err)
			return
		}
		message.StatusHttpSuccess(c)
		return
	}
	message.StatusHttpError(c, httpStatus, err)
	return
}

// Изменить пользователя
func (h *HandlerUser) Update(c *gin.Context) {
	var chuser ChangeUser
	if errr := c.ShouldBindJSON(&chuser); errr != nil {
		message.StatusBadRequestDataH(c, errr)
		return
	}
	id, _ := uuid.Parse(c.Param("id"))
	httpStatus, err := h.userService.UserExist(id)
	if httpStatus == 200 {
		httpStatus, err = h.userService.UpdateUserID(id, chuser)
		if err != nil {
			message.StatusHttpError(c, httpStatus, err)
			return
		}
		message.StatusHttpSuccess(c)
		return
	}
	message.StatusHttpError(c, httpStatus, err)
	return
}
