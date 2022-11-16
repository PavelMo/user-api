package repo

import (
	"refactoring/models"
	"refactoring/storage/user"
)

type userRepo struct {
	st *user.UserStorage
}

func NewUserRepo() *userRepo {
	return &userRepo{st: user.NewUserStorage()}
}

func (u *userRepo) GetAllUsers() models.UserList {
	return u.st.GetAllUsers()
}

func (u *userRepo) CreateUser(name, email string) (string, error) {
	return u.st.CreateUser(name, email)
}

func (u *userRepo) GetUser(id string) (models.User, error) {
	return u.st.GetUser(id)
}

func (u *userRepo) UpdateUser(id, newName string) error {
	return u.st.UpdateUser(id, newName)
}

func (u *userRepo) DeleteUser(id string) error {
	return u.st.DeleteUser(id)
}
