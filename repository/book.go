package repository

import (
	"database/sql"
	"quiz-3-ezra-golang-sanbercode/structs"
)

func GetAllBook(db *sql.DB) (result []structs.Book, err error) {
	sql := "SELECT * FROM book"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book structs.Book

		err = rows.Scan(&book.ID, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.Description, &book.ImageURL, &book.ModifiedAt, &book.ModifiedBy, &book.Price, &book.ReleaseYear, &book.Thickness, &book.Title, &book.TotalPage)
		if err != nil {
			return
		}

		result = append(result, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := "INSERT INTO book(id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"

	errs := db.QueryRow(sql, book.ID, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedAt, book.CreatedBy, book.ModifiedAt, book.ModifiedBy)

	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := "UPDATE book SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8, created_at = $9, created_by = $10, modified_at = $11, modified_by = $12 WHERE id = $13"

	errs := db.QueryRow(sql, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedAt, book.CreatedBy, book.ModifiedAt, book.ModifiedBy, book.ID)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM book WHERE id = $1"

	errs := db.QueryRow(sql, book.ID)
	return errs.Err()
}
