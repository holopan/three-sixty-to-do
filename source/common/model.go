package common

import (
	"fmt"
	"net/http"
)

var (
	UID     int
	TodoMap map[int]TodoDetail
)

//AddTodo is function for add Todo
func AddTodo(title string, detail string) TodoDetail {
	todoDetail := TodoDetail{
		ID:     UID,
		Title:  title,
		Detail: detail,
		Status: "pending",
	}
	TodoMap[UID] = todoDetail
	UID = UID + 1

	return todoDetail
}

//GetAllToDo is function for return TodoList
func GetAllToDo() (*[]TodoDetail, error) {
	var todoList = []TodoDetail{}
	for _, value := range TodoMap {
		todoList = append(todoList, value)
	}

	return &todoList, nil
}

//GetToDo is function for return Todo
func GetToDo(key int) (*TodoDetail, error) {
	value := TodoMap[key]
	if value.ID == 0 {
		return nil, fmt.Errorf("id not found")
	}
	return &value, nil
}

//ModifyTodo is function for add Todo
func ModifyTodo(key int, title string, detail string, status string) (*TodoDetail, error) {
	if TodoMap[key].ID == 0 {
		return nil, fmt.Errorf("id not found")
	}
	todoDetail := TodoDetail{
		ID:     key,
		Title:  title,
		Detail: detail,
		Status: status,
	}
	TodoMap[key] = todoDetail

	return &todoDetail, nil
}

//RemoveTodo is function for remove Todo
func RemoveTodo(key int) error {
	delete(TodoMap, key)

	return nil
}

//ErrorInternal .
var ErrorInternal ErrorResponse = ErrorResponse{
	Error:            "internal_server_error",
	ErrorDescription: "internal server error",
	HTTPStatus:       http.StatusInternalServerError,
}

//ErrorTitleEmpty .
var ErrorTitleEmpty ErrorResponse = ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "receive no title",
	HTTPStatus:       http.StatusBadRequest,
}

//ErrorIDEmpty .
var ErrorIDEmpty ErrorResponse = ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "receive no id",
	HTTPStatus:       http.StatusBadRequest,
}

//ErrorStatusEmpty .
var ErrorStatusEmpty ErrorResponse = ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "receive no status",
	HTTPStatus:       http.StatusBadRequest,
}

//ErrorIDInvalid .
var ErrorIDInvalid ErrorResponse = ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "id invalid",
	HTTPStatus:       http.StatusBadRequest,
}

//ErrorStatusInvalid .
var ErrorStatusInvalid ErrorResponse = ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "status invalid",
	HTTPStatus:       http.StatusBadRequest,
}

//ErrorTodoNotFound .
var ErrorTodoNotFound ErrorResponse = ErrorResponse{
	Error:            "invalid_request",
	ErrorDescription: "To Do list not found",
	HTTPStatus:       http.StatusBadRequest,
}
