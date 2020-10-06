package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// pre
		start := time.Now()

		next.ServeHTTP(writer, request)

		// post
		// TODO: hostname setting from env
		host, _ := os.Hostname()
		writer.Header().Set("X-Response-Time", fmt.Sprintf("%d", time.Now().Sub(start)/time.Microsecond))
		writer.Header().Set("X-Server-Name", host)
	})
}
