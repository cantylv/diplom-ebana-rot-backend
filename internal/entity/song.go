package entity

import (
	"time"

	"github.com/satori/uuid"
)

type Song struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Text        Text      `json:"text" db:"text"`
	Link        string    `json:"link" db:"link"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Text struct {
	Couplets []string `json:"couplets" db:"couplets"`
}
