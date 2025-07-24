package middleware

import (
	"GoAuth/dto"
	util "GoAuth/utils"
	"context"
	"fmt"
	"net/http"
)

func UserLoginRequestValidator(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var payload dto.LoginUserRequestDto
		if err := util.ReadJsonBody(r, &payload); err != nil{
			util.WriteJsonErrorResponse(w, "Invalid request body", err)
			return
		}
		if validatorError := util.Validator.Struct(payload); validatorError != nil{
			util.WriteJsonErrorResponse(w, "Validation failed", validatorError)
			return
		}
		fmt.Println("Payload received for Login :", payload)
		ctx := context.WithValue(r.Context(), "payload", payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CreateUserRequestValidator(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDto
		if err := util.ReadJsonBody(r, &payload); err != nil{
			util.WriteJsonErrorResponse(w, "Invalid request body", err)
			return
		}
		if validatorError := util.Validator.Struct(payload); validatorError != nil{
			util.WriteJsonErrorResponse(w, "Validation failed", validatorError)
			return
		}
			fmt.Println("Payload received for User Creation :", payload)
			ctx := context.WithValue(r.Context(), "payload", payload)
		    next.ServeHTTP(w, r.WithContext(ctx))
            next.ServeHTTP(w, r)
	})
}

