package middleware

import (
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func SetAdminMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("x-user-role")
		if role != "ADMIN" {
			rw.WriteHeader(http.StatusUnauthorized)
			_, _ = rw.Write([]byte("Unauthorized"))
			return
		}
		next.ServeHTTP(rw, r)
	}
}
