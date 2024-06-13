package book

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *Store) DeleteBookByID(bookID, userID uuid.UUID) error {
	query:= `DELETE FROM books WHERE id=$1 AND userid=$2`
	result,err:= s.db.Exec(query,bookID,userID)
	if err!=nil{
		return fmt.Errorf("could not delete book:%v",err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not determine rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no book found with the given id and user id")
	}
	return nil
}