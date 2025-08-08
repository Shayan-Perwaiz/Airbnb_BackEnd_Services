package middleware

import (
	dto "GoHotelService/dtos"
	"GoHotelService/utils"
	"context"
	"fmt"
	"net/http"
)


type contexKey string
const HotelPayloadKey contexKey = "hotelPayload" 

func HotelValidator(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := dto.HotelDto{}
		if err := utils.ReadJsonRequestBody(r, &payload); err != nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid Json body")
			return 
		}

		if validationErr := utils.Validate.Struct(payload); validationErr != nil{
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "validation error")
			return 
		}

		fmt.Println("payload received for hotel", payload)
		ctx := context.WithValue(r.Context(), HotelPayloadKey, payload)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}