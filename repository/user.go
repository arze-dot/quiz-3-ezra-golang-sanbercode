package repository

import (
	"database/sql"
	"quiz-3-ezra-golang-sanbercode/structs"
)

func GetAllUser(db *sql.DB) (result []structs.User, err error) {
	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var user structs.User

		err = rows.Scan(&user.ID, &user.CreatedAt, &user.CreatedBy, &user.ModifiedAt, &user.ModifiedBy, &user.Username)
		if err != nil {
			return
		}

		result = append(result, user)
	}

	return
}

func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users(id, username, password, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	errs := db.QueryRow(sql, user.ID, user.Username, user.Password, user.CreatedAt, user.CreatedBy, user.ModifiedAt, user.ModifiedBy)

	return errs.Err()
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	sql := "UPDATE users SET username = $1, password = $2, created_at = $3, created_by = $4, modified_at = $5, modified_by = $6 WHERE id = $7"

	errs := db.QueryRow(sql, user.Username, user.Password, user.CreatedAt, user.CreatedBy, user.ModifiedAt, user.ModifiedBy, user.ID)

	return errs.Err()
}

func DeleteUser(db *sql.DB, user structs.User) (err error) {
	sql := "DELETE FROM users WHERE id = $1"

	errs := db.QueryRow(sql, user.ID)
	return errs.Err()
}

func GetUserByUsername(db *sql.DB, username string) (structs.User, error) {
	var user structs.User
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}
