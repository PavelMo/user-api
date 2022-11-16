package services

import "refactoring/models"

type IUserService interface {
	GetAllUsers() models.UserList
	CreateUser(name, email string) (string, error)
	GetUser(id string) (models.User, error)
	UpdateUser(id, newName string) error
	DeleteUser(id string) error
}
