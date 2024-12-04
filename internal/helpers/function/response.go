package function

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Response(w http.ResponseWriter, payload any, codeStatus int) {
	w.Header().Add("Content-Type", "application/json")
	body, err := json.Marshal(payload)
	if err != nil {
		w.Header().Add("Content-Length", "0")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(codeStatus)
	contentLength, err := w.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Add("Content-Length", strconv.Itoa(contentLength))
}
