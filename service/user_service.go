package service

import (
	"user-management/model"
	"user-management/repository"
)

type UserService interface {
	GetUsers() ([]model.User, error)
	GetUserByName(userName string) (model.User, error)
	GetUserByID(id string) (model.User, error)
	SaveUser(user model.User) (bool, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (u *userService) GetUsers() ([]model.User, error) {
	return u.repo.GetAll()
}

func (u *userService) GetUserByName(userName string) (model.User, error) {
	return u.repo.GetUserByName(userName)
}

func (u *userService) GetUserByID(id string) (model.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *userService) SaveUser(user model.User) (bool, error) {
	return u.repo.SaveNewUser(user)
}
