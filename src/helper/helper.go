package helper

import (
	"net/http"
)

var (
	UNKNOWN_ERROR = "unknown error"
	INVALID_BODY  = "invalid body"
)

// ServeResponse serves a JSON response
func ServeResponse(w http.ResponseWriter, r *http.Request, status int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if body != nil {
		w.Write(body)
	}
}
