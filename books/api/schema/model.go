package schema

import (
	"time"
)

type Book struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"create_at"`
}
