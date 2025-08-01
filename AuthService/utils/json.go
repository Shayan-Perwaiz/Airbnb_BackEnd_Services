package util

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init(){
	Validator = validator.New(validator.WithRequiredStructEnabled())
}

func WriteJsonResponse(w http.ResponseWriter, status int, data any) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadJsonBody(r *http.Request, result any) error{
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)
}

func WriteJsonSuccessResponse(w http.ResponseWriter, message string, data any) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response_data := map[string] any{
		"message" : message,
		"data" : data,
		"success" : true,
	}
	return json.NewEncoder(w).Encode(response_data)
}

func WriteJsonErrorResponse(w http.ResponseWriter, err error, statusCode int) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response_data := map[string] any{
		"error" : err.Error(),
		"success" : false,
	}
	return json.NewEncoder(w).Encode(response_data)
}
