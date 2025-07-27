package utils

import (
	"encoding/json"
	"net/http"
)

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

func WriteJsonErrorResponse(w http.ResponseWriter, data string, success string) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	response_data := map[string] any{
		"data" : data,
		"success" : success,
	}
	return json.NewEncoder(w).Encode(response_data)
}

func WriteJsonSuccessResponse(w http.ResponseWriter, data string, success string) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response_data := map[string] any{
		"data" : data,
		"success" : success,
	}
	return json.NewEncoder(w).Encode(response_data)
}