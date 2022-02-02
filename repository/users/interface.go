package users

import "Restobook/entities"

type UsersInterface interface {
	Register(newUser entities.User) (entities.User, error)
	LoginUser(email, password string) (entities.User, error)
	Get(userId uint) (entities.User, error)
	Update(userId uint, newUser entities.User) (entities.User, error)
	Delete(userId uint) (entities.User, error)
}