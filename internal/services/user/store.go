package user

import (
	"bookIt/internal/models"
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db:db}
}

func scanRowIntoUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.RegNo,
		&user.Phone,
		&user.Gender,
		&user.IsBanned,
		&user.IsAdmin,
		&user.IsVitian,
		&user.IsVerified,
		&user.IsProfileDone,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
