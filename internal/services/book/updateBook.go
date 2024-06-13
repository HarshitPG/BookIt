package book

import (
	"bookIt/internal/types"

	"github.com/google/uuid"
)

func (s *Store) UpdateBook(bookID uuid.UUID,userID uuid.UUID, book types.UpdateBookPayload) error {
	_, err := s.db.Exec(`
		UPDATE books
		SET Slot1 = $1, Slot2 = $2, Slot3 = $3, Slot4 = $4, Slot5 = $5, Slot6 =$6 , Slot7 = $7 , Slot8 = $8 , Slot9 = $9, Slot10 = $10, Slot11 = $11, Slot12 = $12, Slot13 = $13, Slot14 = $14
		WHERE id = $15 AND UserID = $16`,
	    book.Slot1, book.Slot2, book.Slot3, book.Slot4, book.Slot5, book.Slot6, book.Slot7, book.Slot8, book.Slot9, book.Slot10, book.Slot11, book.Slot12, book.Slot13, book.Slot14, bookID, userID)
	if err != nil {
		return err
	}

	return nil
}