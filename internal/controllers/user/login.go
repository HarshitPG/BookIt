package user

import (
	"bookIt/internal/middleware"
	"bookIt/internal/services/auth"
	"bookIt/internal/types"
	"bookIt/internal/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)


type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}
	u, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found invalid email ",err))
		return
	}
	if !auth.Comparepasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found invalid email or password"))
		return
	}
	secert := []byte(os.Getenv("JWTSecret"))
	userIDString := u.ID.String()
	// fmt.Println("userIDString",userIDString)
	// userIDBytes, _ := strconv.Atoi(userIDString)
	// fmt.Println("userIDString",userIDBytes)
	token, err := middleware.CreateJWT(secert, userIDString)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := struct {
		UserID    uuid.UUID    `json:"id"`
		Email     string `json:"email"`
		Token     string `json:"token"`
	  }{
		UserID:     u.ID,
		Email:     u.Email,
		Token:     token,
	  }
	  
	utils.WriteJSON(w, http.StatusOK, response)
}
