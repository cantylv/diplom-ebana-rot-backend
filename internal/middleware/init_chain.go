package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Init инициализирует цепочку middlewares.
func Init(r *mux.Router, logger *zap.Logger) (h http.Handler) {
	h = recoverm(h, logger)
	h = access(h, logger)
	return h
}
