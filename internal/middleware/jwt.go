package middleware

import (
	"bookIt/internal/types"
	"bookIt/internal/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID string) (string, error) {
	expirationString := os.Getenv("JWTEXPINSEC")
	expirationSeconds, _ := strconv.Atoi(expirationString)
	expiration := time.Second * time.Duration(expirationSeconds)

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userID": userID,
		"expiredAt":time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err!=nil{
		return "",err
	}

	return tokenString,nil
}

type contextKey string

const UserKey contextKey ="userID"

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		tokenString:= utils.GetTokenFromRequest(r)
		log.Printf("Token: %s", tokenString)

		token, err:= validateJWT(tokenString)
		if err !=nil{
			log.Printf("failed to validate token(1): %v",err)
			permissionDenied(w)
			return
		}
		if !token.Valid{
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str:= claims["userID"].(string)

		u,err:= store.GetUserByID(str)
		if err!=nil{
			log.Printf("failed to get userid: %v",err)
			permissionDenied(w)
			return
		}

		ctx:= r.Context()
		log.Printf("ctx(1): %s", ctx)
		ctx= context.WithValue(ctx,UserKey,u.ID.String())
		log.Printf("ctx(2): %s", ctx)
		r= r.WithContext(ctx)
		// log.Printf("r(1): %s", r)

		handlerFunc(w,r)
	}
}

func validateJWT(tokenString string)(* jwt.Token, error){
	return jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){
		if _,ok:= token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil, fmt.Errorf("unexpected method:%v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWTSECRET")),nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) string {
	userID, ok := ctx.Value(UserKey).(string)
	if !ok {
		return "nil"
	}

	return userID
}