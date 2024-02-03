package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hculpan/kabbase/pkg/dbbadger"
	"github.com/hculpan/kabbase/pkg/entities"
	"github.com/hculpan/kabbase/pkg/secure"
)

func setRoutes(r *chi.Mux) {
	r.Post("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body to get username and password
		var user entities.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(user.Username) == 0 || len(user.Passkey) == 0 {
			http.Error(w, "invalid username or passkey", http.StatusUnauthorized)
			return
		}

		data, err := dbbadger.FetchKey(db, "user_"+user.Username)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}

		user2 := entities.User{}
		err = json.Unmarshal(data, &user2)
		if err != nil {
			http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
			return
		}

		fmt.Printf("user: %s\n", user.Passkey)
		fmt.Printf("user2: %s\n", user2.Passkey)
		if !user2.ComparePasskey(user.Passkey) {
			http.Error(w, "invalid username or passkey", http.StatusUnauthorized)
			return
		}

		// Generate PASETO token
		token, err := secure.GenerateToken(user, secretKey, 24*time.Hour) // Token valid for 24 hours
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Return the token
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	})

	// Protected Endpoints
	r.Get("/protected-resource", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{ "message":"you are authorized"}`))
		w.WriteHeader(http.StatusOK)
	})

}
