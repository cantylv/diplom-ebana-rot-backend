package function

import (
	"net/http"

	mc "github.com/cantylv/online-music-lib/internal/helpers/my/constant"
	e "github.com/cantylv/online-music-lib/internal/helpers/my/error"
)

func GetCtxRequestID(r *http.Request) (string, error) {
	requestID, ok := r.Context().Value(mc.ContextKey(mc.RequestID)).(string)
	if !ok {
		return "", e.ErrNoRequestIdInContext
	}
	return requestID, nil
}
