package types

import (
	"bookIt/internal/models"

	"github.com/google/uuid"
)

type BookStore interface {
	GetBookByID(id uuid.UUID) (*models.Book, error)
	GetBooks() ([]*models.Book, error)
	CreateBook(userID uuid.UUID,payload CreateBookPayload) error
	UpdateBook( bookID uuid.UUID,userID uuid.UUID, payload UpdateBookPayload) error
	DeleteBookByID(bookID, userID uuid.UUID) error
	// GetBookByCourseCode(coursecode string) (*models.Book, error)
	// GetBookByCourseName(coursename string) (*models.Book, error)
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	CreateUser(models.User) error
}

type CreateBookPayload struct {
	OwnerName      string    `json:"ownername" validate:"required"`
	Email         string    `json:"email" validate:"required"`
	RegNo         string    `json:"regNo" validate:"required"`
	Phone         string    `json:"phonenumber" validate:"required"`
	CourseCode string `json:"coursecode" validate:"required"`
	CourseName string `json:"coursename" validate:"required"`
	Slot1 string `json:"slot1" validate:"required"`
	Slot2 string `json:"slot2"`
	Slot3 string `json:"slot3"`
	Slot4 string `json:"slot4"`
	Slot5 string `json:"slot5"`
	Slot6 string `json:"slot6"`
	Slot7 string `json:"slot7"`
	Slot8 string `json:"slot8"`
	Slot9 string `json:"slot9"`
	Slot10 string `json:"slot10"`
	Slot11 string `json:"slot11"`
	Slot12 string `json:"slot12"`
	Slot13 string `json:"slot13"`
	Slot14 string `json:"slot14"`
	BookName string `json:"bookname" validate:"required"`
	BookAuthor string `json:"bookauthor" validate:"required"`
	BookImage string `json:"bookimage"`
}


type UpdateBookPayload struct {
	Slot1 string `json:"slot1" validate:"required"`
	Slot2 string `json:"slot2"`
	Slot3 string `json:"slot3"`
	Slot4 string `json:"slot4"`
	Slot5 string `json:"slot5"`
	Slot6 string `json:"slot6"`
	Slot7 string `json:"slot7"`
	Slot8 string `json:"slot8"`
	Slot9 string `json:"slot9"`
	Slot10 string `json:"slot10"`
	Slot11 string `json:"slot11"`
	Slot12 string `json:"slot12"`
	Slot13 string `json:"slot13"`
	Slot14 string `json:"slot14"`
}
