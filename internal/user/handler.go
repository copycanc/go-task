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
}

// Создать нового пользователя
func (h *HandlerUser) Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		message.StatusBadRequestDataH(c, err)
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
}

// Получить пользователя по ID
func (h *HandlerUser) Get(c *gin.Context) {
	id, ok := parseUUIDParam(c, "id")
	if !ok {
		return
	}
	if !h.ensureUserExists(c, id) {
		return
	}
	user, httpStatus, err := h.userService.GetUserID(id)
	if err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	c.JSONP(httpStatus, gin.H{
		"status": "ok",
		"user":   user,
	})
}

// Удалить пользователя по ID
func (h *HandlerUser) Delete(c *gin.Context) {
	id, ok := parseUUIDParam(c, "id")
	if !ok {
		return
	}
	if !h.ensureUserExists(c, id) {
		return
	}
	httpStatus, err := h.userService.DeleteUserID(id)
	if err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	message.StatusHttpSuccess(c)
}

// Изменить пользователя
func (h *HandlerUser) Update(c *gin.Context) {
	var chuser ChangeUser
	if errr := c.ShouldBindJSON(&chuser); errr != nil {
		message.StatusBadRequestDataH(c, errr)
		return
	}
	id, ok := parseUUIDParam(c, "id")
	if !ok {
		return
	}
	if !h.ensureUserExists(c, id) {
		return
	}
	httpStatus, err := h.userService.UpdateUserID(id, chuser)
	if err != nil {
		message.StatusHttpError(c, httpStatus, err)
		return
	}
	message.StatusHttpSuccess(c)
}

func parseUUIDParam(c *gin.Context, name string) (uuid.UUID, bool) {
	id, err := uuid.Parse(c.Param(name))
	if err != nil {
		message.StatusBadRequestDataH(c, err)
		return uuid.Nil, false
	}
	return id, true
}

func (h *HandlerUser) ensureUserExists(c *gin.Context, id uuid.UUID) bool {
	httpStatus, err := h.userService.UserExist(id)
	if httpStatus != 200 {
		message.StatusHttpError(c, httpStatus, err)
		return false
	}
	return true
}
