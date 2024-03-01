package renderer

import (
	"encoding/json"
	"net/http"
)

func RenderResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func RenderError(w http.ResponseWriter, r *http.Request, status int, message string) {
	RenderResponse(w, r, status, map[string]string{"error": message})
}
