package entity

import (
	"time"
)

// Song represents a music song
// @Description Song object
type Song struct {
	ID          string    `json:"id" db:"id" example:"41b3b583-484b-4a49-b683-aece6b539425"`
	Name        string    `json:"name" db:"name" example:"Shape of You"`
	ReleaseDate time.Time `json:"release_date" db:"release_date" example:"01-06-2017"`
	Text        Text      `json:"text" db:"text"` // Text object containing
	Link        string    `json:"link" db:"link" example:"https://youtube.com/YSf231sfsf9"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" example:"01-06-2017"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" example:"01-06-2017"`
}

// Text represents the text of the song.
// @Description Text object containing couplets
type Text struct {
	Couplets []string `json:"couplets" db:"couplets"` // List of couplets in the song
}
