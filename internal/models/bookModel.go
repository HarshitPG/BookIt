package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID `json:"id"`
	UserID        string `json:"userID"`
	OwnerName      string    `json:"ownername"`
	Email         string    `json:"email"`
	RegNo         string    `json:"regNo"`
	Phone         string    `json:"phonenumber"`
	CreatedAt     time.Time `json:"createdAt"`
	CourseCode string `json:"coursecode"`
	CourseName string `json:"coursename"`
	Slot1 string `json:"slot1"`
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
	BookName string `json:"bookname"`
	BookAuthor string `json:"bookauthor"`
	BookImage string `json:"bookimage"`
}



