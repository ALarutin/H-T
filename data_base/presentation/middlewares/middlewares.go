package middlewares

import (
	"data_base/presentation/logger"
	"net/http"
	"time"
)

func MiddlewareLogger(this http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		this.ServeHTTP(w, r)
		logger.Info.Printf("\nmethod: [%s]\naddr: %s\npath: %s\nstart time: %s\n",
			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func MiddlewarePanic(this http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Fatal.Println("recovered", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		this.ServeHTTP(w, r)
	})
}