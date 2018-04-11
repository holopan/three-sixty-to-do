package handler

import (
	"net/http"
	"to-do-list/source/common"
)

//AddToDo is handler function for Router /add
func AddToDo(res http.ResponseWriter, req *http.Request) {
	title := req.FormValue("title")
	detail := req.FormValue("detail")

	writer := common.Writer{
		Resp: res,
	}

	if title == "" {
		writer.ResponseError(common.ErrorTitleEmpty)
		return
	}

	data := common.AddTodo(title, detail)

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})

}
