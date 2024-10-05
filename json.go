package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}
	type errResponse struct {
		// this is using a 'json reflect tag'. It tells json.Marshal() to marshal the Error attribute in go to
		// the 'error' key in a json object. It works in reverse, too, with the json.Unmarshal() method.
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{Error: msg})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
