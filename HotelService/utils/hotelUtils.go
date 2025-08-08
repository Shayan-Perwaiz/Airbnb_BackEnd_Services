package utils

import (
	dto "GoHotelService/dtos"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate := validator.New()
}

func WriteJsonResponse(w http.ResponseWriter, data any) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(data)
}

func ReadJsonRequestBody(r *http.Request, payload *dto.HotelDto) error{
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(payload)
}

func WriteJsonSuccessResponse(w http.ResponseWriter, statusCode int, message string) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response_data := map[string] any{
		"data" : message,
		"success" : true,
	}
	return json.NewEncoder(w).Encode(response_data)

}

func WriteJsonErrorResponse(w http.ResponseWriter, statusCode int, message string) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response_data := map[string] any{
		"data" : message,
		"success" : false,
	}
	return json.NewEncoder(w).Encode(response_data)

}