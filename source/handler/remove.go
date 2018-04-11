package handler

import (
	"net/http"
	"strconv"
	"to-do-list/source/common"
)

//RemoveToDo is handler function for Router /remove
func RemoveToDo(res http.ResponseWriter, req *http.Request) {
	idStr := req.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		//return error
		return
	}

	writer := common.Writer{
		Resp: res,
	}
	data, err := common.GetToDo(id)
	if err != nil {
		//return error
	}
	common.RemoveTodo(id)

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})

}
