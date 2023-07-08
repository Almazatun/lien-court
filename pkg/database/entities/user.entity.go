package entities

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Pass      string    `json:"password"`
	CreatedAt any       `json:"created_at"`
}
