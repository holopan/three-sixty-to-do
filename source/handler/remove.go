package handler

import (
	"net/http"
	"strconv"
	"to-do-list/source/common"
)

//RemoveToDo is handler function for Router /remove
func RemoveToDo(res http.ResponseWriter, req *http.Request) {
	idStr := req.FormValue("id")
	writer := common.Writer{
		Resp: res,
	}

	if idStr == "" {
		writer.ResponseError(common.ErrorIDEmpty)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writer.ResponseError(common.ErrorIDInvalid)
		return
	}

	data, err := common.GetToDo(id)
	if err != nil {
		writer.ResponseError(common.ErrorTodoNotFound)
		return
	}

	err = common.RemoveTodo(id)
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
