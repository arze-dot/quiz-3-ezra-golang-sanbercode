package repository

import (
	"database/sql"
	"quiz-3-ezra-golang-sanbercode/structs"
)

func GetAllCategory(db *sql.DB) (result []structs.Category, err error) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var category structs.Category

		err = rows.Scan(&category.ID, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy, &category.Name)
		if err != nil {
			return
		}

		result = append(result, category)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category(id, name, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5, $6)"

	errs := db.QueryRow(sql, category.ID, category.Name, category.CreatedAt, category.CreatedBy, category.ModifiedAt, category.ModifiedBy)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1, created_at = $2, created_by = $3, modified_at = $4, modified_by = $5 WHERE id = $6"

	errs := db.QueryRow(sql, category.Name, category.CreatedAt, category.CreatedBy, category.ModifiedAt, category.ModifiedBy, category.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)
	return errs.Err()
}
