package books

import (
	"bookIt/internal/middleware"
	"bookIt/internal/types"
	"bookIt/internal/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Handler struct {
	Store types.BookStore
}

func NewHandler(Store types.BookStore) *Handler {
	return &Handler{Store: Store}
}

func (h *Handler) HandleGetBooks(w http.ResponseWriter, r *http.Request){
	books, err:= h.Store.GetBooks()
	if err!=nil{
		utils.WriteError(w, http.StatusInternalServerError,err)
		return
	}
	utils.WriteJSON(w,http.StatusOK, books)
}

func (h *Handler) HandleGetBook(w http.ResponseWriter, r *http.Request){
	bookIdParam := chi.URLParam(r,"id")
	if bookIdParam==""{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("Missing bookId"))
		return
	}
    bookId,err:= uuid.Parse(bookIdParam)
	if err!=nil{
		utils.WriteJSON(w,http.StatusBadRequest,fmt.Errorf("Invalid bookId"))
	}
	book, err := h.Store.GetBookByID(bookId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, book)
}


func (h *Handler) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	var book types.CreateBookPayload
	if err := utils.ParseJSON(r, &book); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(book); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	userID := middleware.GetUserIDFromContext(r.Context())
	if userID == "nil" {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied(1)"))
		return
	}

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not parse userID"))
		return
	}

	err = h.Store.CreateBook(parsedUserID, book)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, book)
}


func (h *Handler) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	var book types.UpdateBookPayload
	if err := utils.ParseJSON(r, &book); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(book); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	bookIdParam := chi.URLParam(r,"id")
	if bookIdParam==""{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("Missing bookId"))
		return
	}
    bookId,err:= uuid.Parse(bookIdParam)
	if err!=nil{
		utils.WriteJSON(w,http.StatusBadRequest,fmt.Errorf("Invalid bookId"))
	}
	userID := middleware.GetUserIDFromContext(r.Context())
	log.Println("userID:",userID)
	if userID == "nil" {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
		return
	}

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not parse userID"))
		return
	}

	err = h.Store.UpdateBook(bookId,parsedUserID,book)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, book)
}

func (h *Handler) HandleDeleteBook(w http.ResponseWriter, r *http.Request){
	bookIdParam := chi.URLParam(r,"id")
	log.Println("bookIdParam:",bookIdParam)

	if bookIdParam==""{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("Missing bookId"))
		return
	}

	bookID,err:= uuid.Parse(bookIdParam)
	if err!=nil{
		utils.WriteJSON(w,http.StatusBadRequest,fmt.Errorf("Invalid bookId"))
	}
	log.Println("bookID:",bookID)
	userID := middleware.GetUserIDFromContext(r.Context())
	log.Println("userID:",userID)
	if userID == "nil" {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
		return
	}

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not parse userID"))
		return
	}

	err = h.Store.DeleteBookByID(bookID, parsedUserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "book deleted successfully"})

}
