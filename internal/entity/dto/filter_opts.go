package dto

type FilterLibraryOptions struct {
	SongIDs         string `json:"song_ids"`
	SongNames       string `json:"song_names"`
	FromReleaseDate string `json:"from_release_date"`
	ToReleaseDate   string `json:"to_release_date"`
	TextPhrases     string `json:"text_phrases"`
	SongLimit       int    `json:"song_limit"`
	SongOffset      int    `json:"song_offset"`
}

type FilterSongOptions struct {
	ID            string `json:"id"`
	CoupletLimit  int    `json:"couplet_limit"`
	CoupletOffset int    `json:"couplet_offset"`
}
