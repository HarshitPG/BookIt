package user

import (
	"bookIt/internal/models"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)


func(s *Store) GetUserByEmail(email string)(*models.User, error){
	rows,err := s.db.Query("SELECT * FROM users WHERE email=$1 ", email)
	if err != nil{
		return nil,err
	}
	u := new(models.User)
	for rows.Next(){
		u,err= scanRowIntoUser(rows)
		if err!=nil{
			return nil,err
		}

	}
	if u.ID ==uuid.Nil {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}


func (s *Store) GetUserByID(id string) (*models.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	u := new(models.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == uuid.Nil {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil

}
