package repo

import "learn/models"

type UserRepo interface {
	Select() ([]models.User, error)
	Insert(u models.User) error
}
