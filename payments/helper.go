package payments

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondCreatedWithLocationHeader(w http.ResponseWriter, url string, id string)  {
	w.Header().Set("Location", fmt.Sprintf("%s/%s", url, id))
	w.WriteHeader(http.StatusCreated)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}