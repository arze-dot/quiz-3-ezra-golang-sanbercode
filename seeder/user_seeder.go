package seeder

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// SeedUsers seeds the database with initial user data
func SeedUsers(db *sql.DB) {
	// User data to seed
	users := []struct {
		ID         int
		Username   string
		Password   string
		CreatedAt  string
		CreatedBy  string
		ModifiedAt string
		ModifiedBy string
	}{
		{
			ID:         1,
			Username:   "admin",
			Password:   "password", // this will be hashed
			CreatedAt:  "2024-08-03T14:55:00Z",
			CreatedBy:  "admin",
			ModifiedAt: "2024-08-03T15:00:00Z",
			ModifiedBy: "admin",
		},
	}

	for _, user := range users {
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Could not hash password: %v", err)
		}

		// Insert user into the database
		_, err = db.Exec(`
			INSERT INTO users (id, username, password, created_at, created_by, modified_at, modified_by)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, user.ID, user.Username, hashedPassword, user.CreatedAt, user.CreatedBy, user.ModifiedAt, user.ModifiedBy)

		if err != nil {
			log.Fatalf("Could not insert user: %v", err)
		}
	}

	log.Println("Users seeded successfully.")
}
