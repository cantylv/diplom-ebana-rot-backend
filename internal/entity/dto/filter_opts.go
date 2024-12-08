package dto

import "time"

type FilterLibraryOptions struct {
	SongIDs         []string `json:"ids" example:"6edba859-b29f-4b27-9203-e6431723b1e4@25539201-1c4f-4f89-894d-3d4aa62b1944"`
	SongNames       []string `json:"names" example:"Take It Slowly@Give Me Everything@The Heat"`
	FromReleaseDate Date     `json:"from_release_date" example:"12-01-2024"`
	ToReleaseDate   Date     `json:"to_release_date" example:"31-12-2024"`
	TextPhrases     string   `json:"text" example:"love"`
	SongLimit       int      `json:"limit" example:"5"`
	SongOffset      int      `json:"offset" example:"2"`
}

type Date struct {
	Valid bool      `json:"valid" example:"false"`
	Time  time.Time `json:"time" swaggertype:"string" example:"2024-12-08"`
}

type FilterSongOptions struct {
	SongID        string `json:"id"`
	CoupletLimit  int    `json:"couplet_limit"`
	CoupletOffset int    `json:"couplet_offset"`
}
