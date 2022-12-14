package repo

import "github.com/PavelMo/user-api/models"

type IUserRepo interface {
	GetAllUsers() models.UserList
	CreateUser(name, email string) (string, error)
	GetUser(id string) (models.User, error)
	UpdateUser(id, newName string) error
	DeleteUser(id string) error
}
