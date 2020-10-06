package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Handle(next http.Handler) http.Handler {
	return AfterSetWrite(next)
}


type responseWriterWithTimer struct {
	http.ResponseWriter

	isHeaderWritten bool
	start           time.Time
}

func AfterSetWrite(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(&responseWriterWithTimer{w, false, time.Now()}, r)
	})
}

func (w *responseWriterWithTimer) WriteHeader(statusCode int) {
	host, _ := os.Hostname()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Response-Time", fmt.Sprintf("%d", time.Now().Sub(w.start)/time.Microsecond))
	w.Header().Set("X-Server-Name", host)

	w.ResponseWriter.WriteHeader(statusCode)
	w.isHeaderWritten = true
}

func (w *responseWriterWithTimer) Write(b []byte) (int, error) {
	if !w.isHeaderWritten {
		w.WriteHeader(200)
	}

	return w.ResponseWriter.Write(b)
}
