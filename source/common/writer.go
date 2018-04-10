package common

import (
	"encoding/json"
	"net/http"
)

// Writer instace of
type Writer struct {
	Resp http.ResponseWriter
}

// ErrorResponse instace of
type ErrorResponse struct {
	Error            string      `json:"error"`
	ErrorDescription string      `json:"error_description"`
	ErrorData        interface{} `json:"error_data,omitempty"`
}

// ResponseError is
func (w Writer) ResponseError(status int, err *ErrorResponse) {
	w.Resp.Header().Set("Content-Type", "application/json")
	w.Resp.WriteHeader(status)

	return
}

// Response is
func (w Writer) Response(data interface{}) {
	w.Resp.Header().Set("Content-Type", "application/json")
	w.Resp.WriteHeader(http.StatusOK)
	if data != nil {
		byteString, _ := json.Marshal(&data)
		w.Resp.Write(byteString)
	}

	return
}
