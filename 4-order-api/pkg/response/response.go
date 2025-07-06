package response

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
}
