package song

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// InitHandlers инициализирует обработчики запросов, отвечающих за работу с песнями
func InitHandlers(r *mux.Router, psqlConn *pgx.Conn, logger *zap.Logger) {
	repoSong := rSong.NewRepoLayer(psqlConn)
	usecaseSong := ucSong.NewUsecaseLayer(repoSong)
	songHandlerManager := song.NewSongHandlerManager(usecaseSong, logger)
	r.HandleFunc("/info", songHandlerManager.GetLibrary).Methods("GET")
	// реализовать CRUD для песен
	r.HandleFunc("/groups/{group_name}/songs/{song_name}", songHandlerManager.DeleteSong).Methods("POST")
	r.HandleFunc("/groups/{group_name}/songs/{song_name}", songHandlerManager.DeleteSong).Methods("GET")
	r.HandleFunc("/groups/{group_name}/songs/{song_name}", songHandlerManager.DeleteSong).Methods("PUT")
	r.HandleFunc("/groups/{group_name}/songs/{song_name}", songHandlerManager.DeleteSong).Methods("DELETE")
}
