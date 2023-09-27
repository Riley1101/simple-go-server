package utils

import (
	"log"
	"net/http"
)

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

type Middleware struct {
	middleware func(http.HandlerFunc) http.HandlerFunc
}

func createMiddleware() Middleware {
	return Middleware{
		middleware: func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				log.Println(r.URL.Path)
				next(w, r)
			}
		},
	}
}
