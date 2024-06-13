package book

import (
	"bookIt/internal/models"
	"database/sql"
)


type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowIntoBooks(rows *sql.Rows) (*models.Book, error) {
	book := new(models.Book)

	err := rows.Scan(
		&book.ID,
		&book.UserID,
		&book.OwnerName,
		&book.Email,
		&book.RegNo,
		&book.Phone,
		&book.CreatedAt,
		&book.CourseCode,
		&book.CourseName,
		&book.Slot1,
		&book.Slot2,
		&book.Slot3,
		&book.Slot4,
		&book.Slot5,
		&book.Slot6,
		&book.Slot7,
		&book.Slot8,
		&book.Slot9,
		&book.Slot10,
		&book.Slot11,
		&book.Slot12,
		&book.Slot13,
		&book.Slot14,
		&book.BookName,
		&book.BookAuthor,
		&book.BookImage,
	)

	if err != nil {
		return nil, err
	}

	return book, nil
}