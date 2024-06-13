package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error{
	if r.Body==nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error{
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int,err error){
	WriteJSON(w,status, map[string]string{"error":err.Error()})
}

func GetTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	// log.Printf("Authorization Header: %s", authHeader)

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		log.Println("User is not authorized or token missing")
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		log.Println("Authorization header format must be Bearer {token}")
		return ""
	}

	// log.Printf("Token Part: %s", parts[1])
	return parts[1]
}
