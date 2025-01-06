package entities

type Message struct {
	ID      int64  `json:"id" gorm:"primaryKey"`
	Message string `json:"Message"`
	Sender  int    `json:"Sender"`
}

// Define the behaviour of message
type MessageService interface {
	SaveMessage(message *Message) error
}
