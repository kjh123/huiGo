package response

import (
    "encoding/json"
    "net/http"
)

// write json response
func WriteJSON(w http.ResponseWriter, v interface{}, code int) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(v)
}

// no content write response
func NoContent(w http.ResponseWriter) {
    w.WriteHeader(http.StatusNoContent)
}

// error response
func Error(w http.ResponseWriter, err string, code int) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(map[string]string{"error": err})
}
