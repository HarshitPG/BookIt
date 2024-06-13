package routes

import (
	"bookIt/internal/controllers/user"

	"github.com/go-chi/chi/v5"
)

type UserRoutes struct {
	handler *user.Handler
}

func NewUserRoutes(handler *user.Handler) *UserRoutes {
	return &UserRoutes{handler: handler}
}

func (u *UserRoutes) RegisterRoutes(router chi.Router) {
	
	router.Post("/login", u.handler.HandleLogin)
	router.Post("/register",  u.handler.HandleRegister)
}
