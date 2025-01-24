package usecases

import (
	"blankproject/entities"
	"blankproject/repository"
)

type MessageService interface {
	Sender(name string) string
	SaveMessage(message *entities.Message) error
	UpdateMessage(message *entities.Message) error
	GetAllMessages() ([]entities.Message, error)
}

type RepoMessageUseCase struct {
	REPO repository.MessageRepository
}

func (g *RepoMessageUseCase) SaveMessage(message *entities.Message) error {
	return g.REPO.Save(message)
}

func (g *RepoMessageUseCase) GetMessageById(id int64) (*entities.Message, error) {
	return g.REPO.GetById(id)
}

func (g *RepoMessageUseCase) GetAllMessages() ([]entities.Message, error) {
	return g.REPO.GetAll()
}
