package usecases

import (
	"blankproject/entities"
	"blankproject/repository"
)

type UserService interface {
	SaveUser(user *entities.User) error
}

type RepoUserUseCase struct {
	REPO repository.UserRepository
}

func (s *RepoUserUseCase) SaveUser(user *entities.User) error {
	return s.REPO.Save(user)
}
