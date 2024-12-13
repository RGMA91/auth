package authenticator

import (
	"auth/security/token"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := verifyToken(w, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Authentication failed")
	}
}

func verifyToken(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return errors.New("missing authorization header")
	}

	tokenString = tokenString[len("Bearer "):]
	err := token.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return err
	}

	log.Println("Valid JWT Token")
	return nil
}
