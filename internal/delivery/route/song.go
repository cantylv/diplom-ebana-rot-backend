package route

import (
	dSong "github.com/cantylv/online-music-lib/internal/delivery/song"
	ucSong "github.com/cantylv/online-music-lib/internal/usecase/song"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// InitHandlers инициализирует обработчики запросов, отвечающих за работу с песнями
func InitHandlers(r *mux.Router, psqlConn *pgx.Conn, logger *zap.Logger) {
	repoSong := rSong.NewRepoLayer(psqlConn)
	usecaseSong := ucSong.NewProccessor()
	songHandlerManager := dSong.NewSongHandlerManager(usecaseSong, logger)

	// реализовать CRUD для песен
	r.HandleFunc("/info", songHandlerManager.GetLibrarySongs).Methods("GET") 
	r.HandleFunc("/groups/{group_name}/songs/{song_name}", songHandlerManager.AddNewSongToLibrary).Methods("POST")
	r.HandleFunc("/groups/{group_name}/songs/{song_name}", songHandlerManager.GetLibrarySong).Methods("GET")
	r.HandleFunc("/groups/{group_name}/songs/{song_name}", songHandlerManager.UpdateLibrarySong).Methods("PUT")
	r.HandleFunc("/groups/{group_name}/songs/{song_name}", songHandlerManager.DeleteLibrarySong).Methods("DELETE")
}
