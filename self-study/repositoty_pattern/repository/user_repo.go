package repository

import models "learn/model"

type UserRepo interface {
	Select() ([]models.User, error)
	Insert(u models.User) error
}
