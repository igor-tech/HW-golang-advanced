package request

import "net/http"

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	body, err := DecodeBody[T](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return body, err
	}

	if err := IsValid(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return body, err
	}

	return body, nil
}
