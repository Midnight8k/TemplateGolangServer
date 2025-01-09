package usecases

import (
	"blankproject/entities"
	"blankproject/repository"
)

type MessageService interface {
	Sender(name string) string
	SaveMessage(message *entities.Message) error
	GetAllMessages() ([]entities.Message, error)
}

type RepoMessageUseCase struct {
	REPO repository.MessageRepository
}

func (s *RepoMessageUseCase) SaveMessage(message *entities.Message) error {
	return s.REPO.Save(message)
}

func (g *RepoMessageUseCase) GetAllMessages() ([]entities.Message, error) {
	return g.REPO.GetAll()
}
