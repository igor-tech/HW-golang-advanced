package middleware

import (
	"net/http"
	"strings"
)

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrapperWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.StatusCode = code
}

func sanitizePath(path string) string {
	safePath := strings.SplitN(path, "?", 2)[0]

	if len(safePath) > 1 && strings.HasSuffix(safePath, "/") {
		safePath = strings.TrimSuffix(safePath, "/")
	}

	return safePath
}
