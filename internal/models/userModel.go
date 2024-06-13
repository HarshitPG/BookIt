package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	RegNo     string    `json:"regno"`
	Phone     string    `json:"phoneNumber"`
	Gender    string    `json:"gender"`
	IsBanned   bool      `json:"-"`
	IsAdmin    bool      `json:"-"`
	IsVitian   bool      `json:"-"`
	IsVerified bool      `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	IsProfileDone bool  `json:"-"`

}

