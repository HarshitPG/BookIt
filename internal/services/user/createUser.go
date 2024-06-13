package user

import (
	"bookIt/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Store) CreateUser(user models.User) error {

	user.ID = uuid.New()

	query := `INSERT INTO users (
		 "id",firstName, lastName, email, password) VALUES (
		$1, $2, $3, $4,$5)`

	_, err := s.db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("could not create user: %v", err)
	}

	return nil
}