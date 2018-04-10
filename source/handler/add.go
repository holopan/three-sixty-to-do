package handler

import (
	"net/http"
	"to-do-list/source/common"
)

//AddToDo is handler function for Router /add
func AddToDo(res http.ResponseWriter, req *http.Request) {
	title := req.FormValue("title")
	detail := req.FormValue("detail")
	size := len(common.TodoList)

	writer := common.Writer{
		Resp: res,
	}
	input := common.TodoDetail{
		ID:     size,
		Title:  title,
		Detail: detail,
		Status: false,
	}
	common.TodoList = append(common.TodoList, input)

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    input,
	})

}
