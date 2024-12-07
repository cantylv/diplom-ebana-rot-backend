package song

import (
	"github.com/cantylv/online-music-lib/internal/entity"
	"github.com/cantylv/online-music-lib/internal/repo/song"
)

type Contract interface {
	GetLibrarySongs() []entity.Song
	AddNewSongToLibrary() entity.Song
	GetLibrarySong() entity.Song
	UpdateLibrarySong() entity.Song
	DeleteLibrarySong() entity.Song
}

var _ Contract = (*proccessor)(nil)

type proccessor struct {
	repoSong song.DBContract
}

func Newproccessor(repoSong song.DBContract) *proccessor {
	return &proccessor{
		repoSong: repoSong,
	}
}

func (t *proccessor) GetLibrarySongs() []entity.Song { return nil }

func (t *proccessor) AddNewSongToLibrary() entity.Song { return entity.Song{} }

func (t *proccessor) GetLibrarySong() entity.Song { return entity.Song{} }

func (t *proccessor) UpdateLibrarySong() entity.Song { return entity.Song{} }

func (t *proccessor) DeleteLibrarySong() entity.Song { return entity.Song{} }
