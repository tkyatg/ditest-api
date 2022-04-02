package usecase

import (
	repo "github.com/tkyatg/ditest-api/repository"
)

type (
	usecase struct {
		repo repo.Repository
	}
	Usecase interface {
		Hello() string
	}
)

func NewUsecase(repo repo.Repository) Usecase {
	return &usecase{
		repo,
	}
}

func (t *usecase) Hello() string {
	return t.repo.Hello()
}
