package logging

import (
	"net/http"
	"time"
)

//Acts as middleware and logs requests received
func LogReceivedThenHandle(f http.HandlerFunc, title string,function string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		end := time.Now()
		AccessReceived("[%s] %q %v", r.Method, r.URL.String(), end.Sub(start))
		
	}
}