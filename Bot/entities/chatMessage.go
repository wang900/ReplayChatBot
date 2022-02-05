package entities

import (
	"time"
)

type ChatMessage struct {
	ID          int       `json:"id"`
	AuthorID    string    `json:"author_id"`
	Message     string    `json:"message"`
	Attachments []string  `json:"attachments"`
	TimeStamp   time.Time `json:"timestamp"`
}
