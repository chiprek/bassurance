package store

import (
	"database/sql"
	"fmt"
)

// CreateUser inserts a new user into the database and returns the generated ID
func CreateUser(db *sql.DB, username, passwordHash, name, role string) (int, error) {
	query := `INSERT INTO users(username, password_hash, name, role)
	VALUES(?,?,?,?)
	RETURNING id;
	`
	var newID int
	err := db.QueryRow(query, username, passwordHash, name, role).Scan(&newID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert user %s: %w", username, err)
	}
	return newID, nil
}
