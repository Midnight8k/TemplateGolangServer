package usecases

import (
	"blankproject/entities"
	"blankproject/repository"
)

// Interface for user Service.
type UserService interface {
	SaveUser(user *entities.User) error
}

// Instantiation of REPO property that will be used to access to the repository of users.
type RepoUserUseCase struct {
	REPO repository.UserRepository
}

// This is the usecase SaveUser.
// This function receives the entity object User, that contains all properties of the user.
// Return: save the entity User in database.
func (s *RepoUserUseCase) SaveUser(user *entities.User) error {
	return s.REPO.Save(user)
}
