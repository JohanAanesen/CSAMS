package logrequest

import (
	"log"
	"net/http"
	"time"
)

// Handler logs HTTP requests
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(time.Now().Format("2006-01-02 03:04:05"), r.RemoteAddr, r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}
