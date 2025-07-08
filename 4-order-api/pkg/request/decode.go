package request

import (
	"encoding/json"
	"net/http"
)

func DecodeBody[T any](r *http.Request) (T, error) {
	var body T
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return body, err
	}
	return body, nil
}
