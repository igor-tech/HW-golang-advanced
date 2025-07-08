package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapperWriter := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapperWriter, r)
		duration := time.Since(start)
		log.WithFields(log.Fields{
			"path":     sanitizePath(r.URL.Path),
			"method":   r.Method,
			"duration": duration,
			"status":   wrapperWriter.StatusCode,
		}).Info("Handled request")
	})
}
