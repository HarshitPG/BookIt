package user

import (
	"bookIt/internal/models"
	"bookIt/internal/services/auth"
	"bookIt/internal/types"
	"bookIt/internal/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)


func (h *Handler) HandleRegister(w http.ResponseWriter,r *http.Request){
	var payload types.RegisterUserPayload
	if err:= utils.ParseJSON(r,&payload); err !=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	if err := utils.Validate.Struct(payload);err!=nil{
		error:=err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("invalid payload %v",error))
		return
	}

	_,err:= h.store.GetUserByEmail(payload.Email)
	if err == nil{
		utils.WriteError(w,http.StatusBadRequest, fmt.Errorf("user with email %s already exist",payload.Email))
	}
	hashedPassword,err:= auth.HashPassword(payload.Password)
	if err!=nil{
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}

	err = h.store.CreateUser(models.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
		IsProfileDone: false,
	})
	if err!=nil{
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}
	utils.WriteJSON(w,http.StatusCreated,"Success")
}