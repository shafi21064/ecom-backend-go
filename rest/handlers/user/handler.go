package user

import "github.com/shafi21064/ecom-go/repo"

type Handler struct{
	userRepo repo.UserRepo
}

func NewHandler(userRepo repo.UserRepo) *Handler{
	return &Handler{
		userRepo: userRepo,
	}
}