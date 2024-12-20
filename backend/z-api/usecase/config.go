package usecase

import (
	"zapi/repository"
)

type Usecase struct {
	Repo *repository.Repository
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		Repo: repo,
	}
}
