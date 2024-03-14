package controllers

import (
	"encoding/json"
	"net/http"

	m "uts/models"
)

func sendSuccessResponse(w http.ResponseWriter, message string) {
	var response m.SuccessResponse
	response.Status = 200
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response m.ErrorResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
