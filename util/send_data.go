package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	log.Print(data)
	encoder.Encode(data)
}

func SendError(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}
