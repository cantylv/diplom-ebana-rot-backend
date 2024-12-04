package middleware

import (
	"fmt"
	"net/http"

	"github.com/cantylv/online-music-lib/internal/entity/dto"
	f "github.com/cantylv/online-music-lib/internal/helpers/function"
	e "github.com/cantylv/online-music-lib/internal/helpers/my/error"
	"go.uber.org/zap"
)

// recoverm миддлвар для обработки паники, возникающей в работе сервера.
// В случае паники возвращается json-объект c сообщением об ошибке внутри сервера и статусом 500.
func recoverm(h http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(fmt.Sprintf("error while handling request: %v", err))
				f.Response(w, dto.ResponseError{Error: e.ErrInternal.Error()}, http.StatusInternalServerError)
				return
			}
		}()
		h.ServeHTTP(w, r)
	})
}
