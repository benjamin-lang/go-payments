package middleware

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func RespondCreatedWithLocationHeader(w http.ResponseWriter, url string, id string) {
    w.Header().Set("Location", fmt.Sprintf("%s/%s", url, id))
    w.WriteHeader(http.StatusCreated)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}
