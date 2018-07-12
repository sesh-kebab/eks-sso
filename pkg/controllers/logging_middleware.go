package controllers

import (
	"log"
	"net/http"
)

// NewLoggingMiddleware returns middleware that logs http request and response
func NewLoggingMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//todo: need to remove this log message
			// session, err := store.Get(r, cookieSession)
			// if err != nil {
			// 	log.Println("[ERROR]", err)
			// } else {
			// 	for k, v := range session.Values {
			// 		log.Println("[DEBUG]", k, v)
			// 	}
			// }

			log.Println("[DEBUG] request", r.RequestURI)
			resp := ResponseLogger{
				ResponseWriter: w,
				Status:         200,
				ResponseBody:   "",
			}

			next.ServeHTTP(&resp, r)

			log.Println("[DEBUG] response", resp.Status)
			if resp.Status >= 400 {
				log.Println("[ERROR]", resp.ResponseBody)
			}
		})
	}
}

// ResponseLogger stores the response body and status code
// which can be used to log the response
type ResponseLogger struct {
	http.ResponseWriter
	Status       int
	ResponseBody string
}

// WriteHeader caches response code
func (r *ResponseLogger) WriteHeader(code int) {
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
}

// Write caches response body
func (r *ResponseLogger) Write(b []byte) (int, error) {
	r.ResponseBody = string(b)
	return r.ResponseWriter.Write(b)
}
