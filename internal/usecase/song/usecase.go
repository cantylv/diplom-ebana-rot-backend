package song

import "github.com/cantylv/online-music-lib/internal/entity"

type Contract interface {
	GetLibrarySongs() []entity.Song
	AddNewSongToLibrary() entity.Song
	GetLibrarySong() entity.Song
	UpdateLibrarySong() entity.Song
	DeleteLibrarySong() entity.Song
}

type Proccessor struct {
	// репо
}

var _ Contract = (*Proccessor)(nil)

func NewProccessor() *Proccessor {
	return &Proccessor{}
}

func (t *Proccessor) GetLibrarySongs() []entity.Song { return nil }

func (t *Proccessor) AddNewSongToLibrary() entity.Song { return entity.Song{} }

func (t *Proccessor) GetLibrarySong() entity.Song { return entity.Song{} }

func (t *Proccessor) UpdateLibrarySong() entity.Song { return entity.Song{} }

func (t *Proccessor) DeleteLibrarySong() entity.Song { return entity.Song{} }
