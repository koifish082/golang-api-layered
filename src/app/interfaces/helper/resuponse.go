package helper

import (
	"encoding/json"
	"net/http"
)

// Succeed creates API success response
func Succeed(w http.ResponseWriter, response interface{}) {
	createResponse(w, http.StatusOK, response)
}

// Fail creates API failure response
func Fail(w http.ResponseWriter, code int, response interface{}) {
	createResponse(w, code, response)
}

func createResponse(w http.ResponseWriter, code int, response interface{}) {
	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8;")
	w.WriteHeader(code)
	w.Write(res)
}
