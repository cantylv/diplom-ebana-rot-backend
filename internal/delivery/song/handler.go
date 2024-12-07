package song

import (
	"net/http"

	"github.com/cantylv/online-music-lib/internal/usecase/song"
	"go.uber.org/zap"
)

type SongHandlerManager struct {
	ucSong song.Contract
	logger *zap.Logger
}

func NewSongHandlerManager(ucSong song.Contract, logger *zap.Logger) *SongHandlerManager {
	return &SongHandlerManager{
		ucSong: ucSong,
		logger: logger,
	}
}

func (t *SongHandlerManager) GetLibrarySongs(w http.ResponseWriter, r *http.Request) {
	
}

func (t *SongHandlerManager) AddNewSongToLibrary(w http.ResponseWriter, r *http.Request) {}

func (t *SongHandlerManager) GetLibrarySong(w http.ResponseWriter, r *http.Request) {}

func (t *SongHandlerManager) UpdateLibrarySong(w http.ResponseWriter, r *http.Request) {}

func (t *SongHandlerManager) DeleteLibrarySong(w http.ResponseWriter, r *http.Request) {}
