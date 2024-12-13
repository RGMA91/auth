package login

import (
	"auth/database"
	"auth/security/token"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserLogin struct {
	Email    string
	Password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u UserLogin
	json.NewDecoder(r.Body).Decode(&u)

	db := database.DatabaseConnection()
	defer db.Close()

	isValid, err := verifyPassword(db, u.Email, u.Password)
	if err == nil && isValid {
		log.Println("Password verified successfully")
		tokenString, err := token.CreateToken(u.Email)
		if err != nil {
			log.Println("Error creating token")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal error")
		}
		log.Println("Token created successfully: " + tokenString)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		log.Println("Password verification failed")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func verifyPassword(db *sql.DB, email, password string) (bool, error) {
	// Retrieve salt and hash for the given email
	var salt, storedHash string
	query := "SELECT salt, passhash FROM account WHERE email = $1"
	err := db.QueryRow(query, email).Scan(&salt, &storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No such account
		}
		return false, fmt.Errorf("failed to query account: %w", err)
	}
	// Hash the provided password with the stored salt
	computedHash := hashPassword(password, salt)

	// Compare hashes
	return computedHash == storedHash, nil
}

func hashPassword(password, salt string) string {
	hash := sha256.Sum256([]byte(salt + password))
	return hex.EncodeToString(hash[:])
}
