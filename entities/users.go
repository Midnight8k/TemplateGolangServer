package entities

// Struct that contains the user parameters.
type User struct {
	ID           int64  `json:"id" gorm:"primaryKey"`
	Nickname     string `json:"Nickname"`
	Username     string `json:"Username"`
	Email        string `json:"Email"`
	Notes        string `json:"Notes"`
	CreationDate int64  `json:"CreateDate"`
}

// Define the behaviour of user
type UserService interface {
	SaveUser(user *User) error
}
