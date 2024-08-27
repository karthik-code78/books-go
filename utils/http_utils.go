package utils

import "net/http"

func SendErrorResponse(res http.ResponseWriter, message string, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	errorResponse, err := JsonMarshaller(message, "error", statusCode)
	if err != nil {
		http.Error(res, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	_, err = res.Write(errorResponse)
	if err != nil {
		http.Error(res, "Error writing JSON response", http.StatusInternalServerError)
	}
}
