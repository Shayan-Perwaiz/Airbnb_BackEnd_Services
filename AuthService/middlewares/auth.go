package middleware

import (
	config "GoAuth/Config/env"
	util "GoAuth/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthoriationMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			util.WriteJsonErrorResponse(w, errors.New("authorization header is required"), http.StatusUnauthorized)
			return 
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			util.WriteJsonErrorResponse(w, errors.New("authorization header must start with Bearer"), http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			util.WriteJsonErrorResponse(w, errors.New("token is required"), http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
			return []byte(config.GetString("JWT_TOKEN", "TOKEN")), nil
		})

		if err != nil{
			util.WriteJsonErrorResponse(w, errors.New("invalid token"), http.StatusUnauthorized)
			return
		}

		userId, okId := claims["user_id"].(float64)
		userEmail, okEmail := claims["user_email"].(string)
		userName, okName := claims["user_name"].(string)

		if !okId || !okEmail || !okName {
			util.WriteJsonErrorResponse(w, errors.New("invalid token claims"), http.StatusUnauthorized)
		}

		fmt.Println("Authenticated user ID:", int64(userId), "Email:", userEmail, "Name:", userName)
		ctx := context.WithValue(r.Context(), "userID", userId)
		ctx = context.WithValue(ctx, "email", userEmail)
		ctx = context.WithValue(ctx, "name", userName)

		next.ServeHTTP(w, r.WithContext(ctx))


	})
}