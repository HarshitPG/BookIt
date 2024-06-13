package book

import (
	"bookIt/internal/types"

	"github.com/google/uuid"
)

func (s *Store) CreateBook( userID uuid.UUID,book types.CreateBookPayload) error {
	bookID := uuid.New()
	_, err := s.db.Exec(`INSERT INTO books ("id",UserID,OwnerName,Email,RegNo,Phone,CourseCode,CourseName,Slot1,Slot2, Slot3, Slot4, Slot5, Slot6, Slot7, Slot8, Slot9, Slot10,Slot11,Slot12,Slot13,Slot14,BookName,BookAuthor,BookImage ) VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9,$10,$11,$12,$13, $14, $15, $16, $17, $18, $19,$20,$21,$22,$23,$24,$25)`, bookID,userID, book.OwnerName, book.Email, book.RegNo, book.Phone, book.CourseCode, book.CourseName, book.Slot1, book.Slot2, book.Slot3, book.Slot4, book.Slot5, book.Slot6, book.Slot7, book.Slot8, book.Slot9, book.Slot10, book.Slot11, book.Slot12, book.Slot13, book.Slot14, book.BookName, book.BookAuthor, book.BookImage)
	if err != nil {
		return err
	}

	return nil
}