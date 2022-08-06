package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, data interface{}, code int) {
	buf, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	status := http.StatusOK
	if code != 0 {
		status = code
	}
	w.WriteHeader(status)
	w.Write(buf)
}

func ErrJSON(w http.ResponseWriter, err error, code int) {
	log.Printf("An error occurred: %s", err.Error())
	JSON(w, map[string]string{
		"error": err.Error(),
	}, code)
}
