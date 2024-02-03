package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/hculpan/kabbase/pkg/secure"
)

// AuthMiddleware to validate PASETO token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.String(), "/authenticate") {
			next.ServeHTTP(w, r)
			return
		}

		auth := r.Header.Get("Authorization")
		if len(auth) == 0 || !strings.HasPrefix(auth, "Bearer ") {
			sendError(w, "missing authorization", http.StatusUnauthorized)
			return
		}
		auth = auth[7:]

		perms, err := secure.ValidateToken(auth, secretKey)
		if err != nil {
			sendError(w, fmt.Sprintf("error during authorization: %s", err), http.StatusUnauthorized)
			return
		} else if len(perms) == 0 {
			sendError(w, "Unknown error during authorization", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("%s requested\n", r.URL.String())
		next.ServeHTTP(w, r)
	})
}

// Other functions for handling authentication, token generation, and BadgerDB interactions
