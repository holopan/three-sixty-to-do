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

	id, err := strconv.Atoi(idStr)
	if err != nil {
		//return error
		return
	}
	if status != "pending" && status != "done" {
		//return error
		return
	}

	writer := common.Writer{
		Resp: res,
	}

	data, err := common.ModifyTodo(id, title, detail, status)
	if err != nil {
		//return error
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

	id, err := strconv.Atoi(idStr)
	if err != nil {
		//return error
		return
	}
	if status != "pending" && status != "done" {
		//return error
		return
	}

	writer := common.Writer{
		Resp: res,
	}

	input, err := common.GetToDo(id)
	if err != nil {
		//return error
		return
	}
	input.Status = status

	data, err := common.ModifyTodo(input.ID, input.Title, input.Detail, input.Status)
	if err != nil {
		//return error
		return
	}

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})

}
