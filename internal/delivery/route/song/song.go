package song

import (
	dSong "github.com/cantylv/online-music-lib/internal/delivery/song"
	rSong "github.com/cantylv/online-music-lib/internal/repo/song"
	ucSong "github.com/cantylv/online-music-lib/internal/usecase/song"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// InitHandlers инициализирует обработчики запросов, отвечающих за работу с песнями
func InitHandlers(r *mux.Router, psqlConn *pgx.Conn, logger *zap.Logger) {
	repoSong := rSong.NewDatabaseConnector(psqlConn)
	usecaseSong := ucSong.Newproccessor(repoSong)
	songHandlerManager := dSong.NewSongHandlerManager(usecaseSong, logger)

	// реализовать CRUD для песен
	r.HandleFunc("/api/v1/songs", songHandlerManager.GetLibrarySongs).Methods("GET")
	r.HandleFunc("/api/v1/songs", songHandlerManager.AddNewSongToLibrary).Methods("POST")
	r.HandleFunc("/api/v1/songs/{song_id}", songHandlerManager.GetLibrarySong).Methods("GET")
	r.HandleFunc("/api/v1/songs/{song_id}", songHandlerManager.UpdateLibrarySong).Methods("PUT")
	r.HandleFunc("/api/v1/songs/{song_id}", songHandlerManager.DeleteLibrarySong).Methods("DELETE")
}
