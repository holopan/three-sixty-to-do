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
	HTTPStatus       int         `json:"-"`
	ErrorData        interface{} `json:"error_data,omitempty"`
}

// SuccessResponse struct of response when success case
type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//TodoDetail is struct for
type TodoDetail struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

// Response is
func (w Writer) Response(data SuccessResponse) {
	w.Resp.Header().Set("Content-Type", "application/json")
	w.Resp.WriteHeader(http.StatusOK)
	if data.Data != nil {
		byteString, _ := json.Marshal(&data)
		w.Resp.Write(byteString)
	}

	return
}

func (w Writer) ResponseError(data ErrorResponse) {
	w.Resp.Header().Set("Content-Type", "application/json")
	w.Resp.WriteHeader(data.HTTPStatus)
	byteString, _ := json.Marshal(&data)
	w.Resp.Write(byteString)

	return
}
