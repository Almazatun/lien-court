package entities

import "github.com/google/uuid"

type Link struct {
	ID        uuid.UUID `json:"id"`
	Short     string    `json:"short"`
	Original  string    `json:"original"`
	CreatedAt any       `json:"created_at"`
}

type Links struct {
	Links []Link `json:"links"`
}
