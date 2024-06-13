package main

import (
	"bookIt/internal/controllers/books"
	"bookIt/internal/controllers/user"
	"bookIt/internal/db"
	routesb "bookIt/internal/routes/book"
	routesu "bookIt/internal/routes/user"
	servicesb "bookIt/internal/services/book"
	servicesu "bookIt/internal/services/user"
	"log"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct{
	addr string
	db *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}


func(s *APIServer) Run() error{
	router:= chi.NewRouter()
	router.Use(middleware.Logger)
	subrouter := router.Route("/api/v1", func(r chi.Router) {})
	
	userStore := servicesu.NewStore(s.db) 
	userHandler := user.NewHandler(userStore) 
	userRoutes := routesu.NewUserRoutes(userHandler)
	userRoutes.RegisterRoutes(subrouter) 


	bookStore := servicesb.NewStore(s.db) 
	bookHandler := books.NewHandler(bookStore) 
	bookRoutes := routesb.NewBookRoutes(bookHandler,userStore)
	bookRoutes.RegisterRoutes(subrouter) 

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connected")
}

func main() {
	config := db.Envs

	db, err := db.NewSQLStorage(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	initStorage(db)

	server := NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
