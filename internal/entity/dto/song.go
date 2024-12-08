package dto

// CreateData represents the data required to create a new song.
// @Description Structure for creating a new song with its details.
type CreateData struct {
	Name        string  `json:"name" db:"name" example:"Shape of You"`
	ReleaseDate string  `json:"release_date" db:"release_date" example:"01-06-2017"`
	Text        NewText `json:"text" db:"text"` // Text object containing
	Link        string  `json:"link" db:"link" example:"https://youtube.com/YSf231sfsf9"`
}

// UpdateSong represents the data required to update an existing song.
// @Description Structure for updating an existing song with its details.
type UpdateSong struct {
	ID          string  `json:"id" example:"41b3b583-484b-4a49-b683-aece6b539425"`
	Name        string  `json:"name" db:"name" example:"Shape of You"`
	ReleaseDate string  `json:"release_date" db:"release_date" example:"01-06-2017"`
	NewText     NewText `json:"text" db:"text"` // Text object containing
	Link        string  `json:"link" db:"link" example:"https://youtube.com/YSf231sfsf9"`
}

// NewText represents the text of the song, including its couplets.
// @Description Structure for the text of the song, including its couplets.
type NewText struct {
	Couplets []string `json:"couplets"` // List of couplets in the song
}
