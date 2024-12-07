package dto

import (
	"time"

	"github.com/satori/uuid"
)

type CreateData struct {
	Name        string    `json:"name" db:"name"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Text        NewText   `json:"text" db:"new_text"`
	Link        string    `json:"link" db:"link"`
}

type UpdateSong struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"release_date"`
	NewText     NewText   `json:"new_text"`
	Link        string    `json:"link"`
}

type NewText struct {
	Couplets []string `json:"couplets"`
}
