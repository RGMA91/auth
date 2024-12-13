package register

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"auth/database"

	_ "github.com/lib/pq"
)

const saltLength = 128

type UserRegistrationData struct {
	Username string
	Email    string
	Password string
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	db := database.DatabaseConnection()
	defer db.Close()

	var user UserRegistrationData
	json.NewDecoder(r.Body).Decode(&user)

	err := createAccount(db, user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("Failed to create account: %v", err)
	} else {
		log.Println("Account created successfully")
	}
}

func createAccount(db *sql.DB, username, email, password string) error {
	// Generate salt
	salt, err := generateSalt(saltLength)
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	// Hash password with salt
	passHash := hashPassword(password, salt)

	// Insert account into database
	query := "INSERT INTO account (username, email, salt, passhash) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(query, username, email, salt, passHash)
	if err != nil {
		return fmt.Errorf("failed to insert account: %w", err)
	}

	return nil
}

func generateSalt(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func hashPassword(password, salt string) string {
	hash := sha256.Sum256([]byte(salt + password))
	return hex.EncodeToString(hash[:])
}
