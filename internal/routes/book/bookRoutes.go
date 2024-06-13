package routes

import (
	"bookIt/internal/controllers/books"
	"bookIt/internal/middleware"
	"bookIt/internal/types"

	"github.com/go-chi/chi/v5"
)

type BookRoutes struct {
	handler *books.Handler
	userStore types.UserStore
}

func NewBookRoutes(handler *books.Handler, userStore types.UserStore) *BookRoutes{
	return &BookRoutes{handler: handler ,userStore: userStore}
}

func (b *BookRoutes) RegisterRoutes(router chi.Router) {
	
	router.Get("/books", b.handler.HandleGetBooks)
	router.Get("/book/{id}",  b.handler.HandleGetBook)
	router.Post("/book", middleware.WithJWTAuth(b.handler.HandleCreateBook, b.userStore))
	router.Delete("/book/{id}", middleware.WithJWTAuth(b.handler.HandleDeleteBook, b.userStore))
	router.Put("/book/{id}",middleware.WithJWTAuth(b.handler.HandleUpdateBook, b.userStore))
}
