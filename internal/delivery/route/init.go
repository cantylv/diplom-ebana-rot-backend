package route

import (
	"net/http"

	"github.com/cantylv/online-music-lib/internal/delivery/route/song"
	"github.com/cantylv/online-music-lib/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// InitHTTPHandlers инициализирует обработчики запросов, а также добавляет цепочку middlewares в обработку запроса.
func InitHTTPHandlers(r *mux.Router, psqlConn *pgx.Conn, logger *zap.Logger) http.Handler {
	subrouter := r.PathPrefix("/api/v1").Subrouter()
	song.InitHandlers(subrouter, psqlConn, logger)
	return middleware.Init(subrouter, logger)
}
