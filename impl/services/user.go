package services

import (
	"refactoring/interfaces/repo"
	"refactoring/models"
)

type userService struct {
	repo repo.IUserRepo
}

func NewUserService(repo repo.IUserRepo) *userService {
	return &userService{repo: repo}
}

func (u *userService) GetAllUsers() models.UserList {
	return u.repo.GetAllUsers()
}

func (u *userService) CreateUser(name, email string) (string, error) {
	return u.repo.CreateUser(name, email)
}

func (u *userService) GetUser(id string) (models.User, error) {
	return u.repo.GetUser(id)
}

func (u *userService) UpdateUser(id, newName string) error {
	return u.repo.UpdateUser(id, newName)
}

func (u *userService) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}
