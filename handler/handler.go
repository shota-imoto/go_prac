package handler

import (
	"net/http"
)

func Handle(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f.ServeHTTP(w, r)
	}
}
