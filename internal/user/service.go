package user

import (
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type UserService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (u *UserService) GetAllUser() (map[uuid.UUID]UserOutput, int, error) {
	userWithoutPass := make(map[uuid.UUID]UserOutput)
	users, err := u.storage.GetAllUser()
	if err != nil {
		slog.Error("STORAGE: get user failed", "err", err)
		return nil, 500, errors.New("ошибка при получении данных")
	}
	for id, user := range users {
		userWithoutPass[id] = user.OutputUser()
	}
	return userWithoutPass, 200, nil
}

func (u *UserService) EmailExist(email string) (int, error) {
	exist, err := u.storage.ExistEmailUser(email)
	if err != nil {
		slog.Error("STORAGE: get user failed", "err", err)
		return 500, errors.New("ошибка при получении данных")
	}
	if exist {
		return 400, errors.New("пользователь с данным Email уже зарегистрирован")
	}
	return 200, nil
}

func (u *UserService) CreateUser(user User) (int, error) {
	user = User{
		ID:       uuid.New(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := u.storage.SaveUser(user); err != nil {
		slog.Error("STORAGE: save user failed", "err", err)
		return 500, errors.New("ошибка при сохранении")
	}
	return 200, nil
}

func (u *UserService) UserExist(uuid uuid.UUID) (int, error) {
	exist, err := u.storage.ExistUser(uuid)
	if err != nil {
		slog.Error("STORAGE: get user failed", "err", err)
		return 500, errors.New("ошибка при получении данных")
	}
	if !exist {
		return 404, errors.New("пользователь не найден")
	}
	return 200, nil
}

func (u *UserService) GetUserID(uuid uuid.UUID) (*UserOutput, int, error) {
	user, err := u.storage.GetUserID(uuid)
	if err != nil {
		slog.Error("STORAGE: get user failed", "err", err)
		return nil, 500, errors.New("ошибка при получении данных")
	}
	userOutputPtr := user.OutputUser()
	return &userOutputPtr, 200, nil
}

func (u *UserService) DeleteUserID(uuid uuid.UUID) (int, error) {
	if err := u.storage.DeleteUser(uuid); err != nil {
		slog.Error("STORAGE: delete user failed", "err", err)
		return 500, errors.New("ошибка при удалении данных")
	}
	return 200, nil
}

func (u *UserService) UpdateUserID(uuid uuid.UUID, chuser ChangeUser) (int, error) {
	user, err := u.storage.GetUserID(uuid)
	if err != nil {
		slog.Error("STORAGE: get user failed", "err", err)
		return 500, errors.New("ошибка при получении данных")
	}
	if (chuser.NewPassword != "" && chuser.OldPassword == "") || (chuser.NewPassword == "" && chuser.OldPassword != "") {
		return 400, errors.New("при смене пароля необходимо ввести старый и новый пароль")
	}
	if ChekChangeEmail(chuser) {
		exist, erre := u.storage.ExistEmailUser(chuser.Email)
		if erre != nil {
			slog.Error("STORAGE: get user failed", "err", erre)
			return 500, errors.New("ошибка при получении данных")
		}
		if exist {
			return 400, errors.New("пользователь с данным Email уже зарегистрирован")
		}
		user.Email = chuser.Email
	}
	if ChekChangePass(chuser) {
		if user.Password != chuser.OldPassword {
			return 400, errors.New("старый пароль не совпадает с введенным")
		}
		user.Password = chuser.NewPassword
	}
	if errs := u.storage.SaveUser(*user); errs != nil {
		slog.Error("STORAGE: save user failed", "err", errs)
		return 500, errors.New("ошибка при сохранении")
	}
	return 200, nil
}

func ChekChangePass(u ChangeUser) bool {
	return u.NewPassword != "" && u.OldPassword != ""
}

func ChekChangeEmail(u ChangeUser) bool {
	return u.Email != ""
}
