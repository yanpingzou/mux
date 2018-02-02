package httputils

import (
	"encoding/json"
	"net/http"
)

// Response response the request.
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// WriteJSON writes the value v to the http response stream as json with standard json encoding.
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// MakeResponse make success http response.
func MakeResponse(w http.ResponseWriter, code int, v interface{}) error {
	response := Response{
		Status:  code,
		Message: "success",
		Data:    v,
	}
	return WriteJSON(w, code, response)
}

// MakeErrResponse make error http response.
func MakeErrResponse(w http.ResponseWriter, err error) error {
	statusCode := GetHTTPErrorStatusCode(err)
	response := Response{
		Status:  statusCode,
		Message: err.Error(),
	}
	return WriteJSON(w, statusCode, response)
}
