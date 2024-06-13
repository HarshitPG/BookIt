package book

import (
	"bookIt/internal/models"

	"github.com/google/uuid"
)




func (s *Store) GetBooks()([]*models.Book,error){
	rows,err := s.db.Query("SELECT * FROM books")
	if err!=nil{
		return nil, err
	}
	books := make([]*models.Book, 0)
	for rows.Next(){
		p,err := scanRowIntoBooks(rows)
		if err!= nil{
			return nil,err
		}
		books= append(books,p)
	}
	return books, nil

}

func (s *Store) GetBookByID(bookId uuid.UUID) ( *models.Book, error){
	rows,err := s.db.Query("SELECT * FROM books WHERE id =$1", bookId)
	if err != nil {
		return nil, err
	}
	p := new(models.Book)
	for rows.Next() {
		p, err = scanRowIntoBooks(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil

}






