package handler

import (
	"net/http"
	"strconv"
	"to-do-list/source/common"
)

//ModifyToDo is handler function for Router /add
func ModifyToDo(res http.ResponseWriter, req *http.Request) {
	idStr := req.FormValue("id")
	title := req.FormValue("title")
	detail := req.FormValue("detail")
	status := req.FormValue("status")

	writer := common.Writer{
		Resp: res,
	}

	if idStr == "" {
		writer.ResponseError(common.ErrorIDEmpty)
		return
	}
	if title == "" {
		writer.ResponseError(common.ErrorTitleEmpty)
		return
	}
	if status == "" {
		writer.ResponseError(common.ErrorStatusEmpty)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writer.ResponseError(common.ErrorIDInvalid)
		return
	}

	if status != "pending" && status != "done" {
		writer.ResponseError(common.ErrorStatusInvalid)
		return
	}

	data, err := common.ModifyTodo(id, title, detail, status)
	if err != nil {
		writer.ResponseError(common.ErrorInternal)
		return
	}

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})

}

//ModifyStatusToDo is handler function for Router /add
func ModifyStatusToDo(res http.ResponseWriter, req *http.Request) {
	idStr := req.FormValue("id")
	status := req.FormValue("status")

	writer := common.Writer{
		Resp: res,
	}

	if idStr == "" {
		writer.ResponseError(common.ErrorIDEmpty)
		return
	}
	if status == "" {
		writer.ResponseError(common.ErrorStatusEmpty)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writer.ResponseError(common.ErrorIDInvalid)
		return
	}

	input, err := common.GetToDo(id)
	if err != nil {
		writer.ResponseError(common.ErrorTodoNotFound)
		return
	}

	input.Status = status

	data, err := common.ModifyTodo(input.ID, input.Title, input.Detail, input.Status)
	if err != nil {
		writer.ResponseError(common.ErrorInternal)
		return
	}

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})

}
