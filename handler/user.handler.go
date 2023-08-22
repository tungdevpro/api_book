package handler

import "api_book/repository"

type UserHandler struct {
	UserRepository repository.UserRepository
}

func (user *UserHandler) GetProfile() {}
