package handler

import (
	"net/http"
	"strconv"
	"strings"
	"to-do-list/source/common"
)

//ViewHandler is handler function for manage Router /view/
func ViewHandler(res http.ResponseWriter, req *http.Request) {

	path := strings.TrimPrefix(req.URL.Path, "/view/")
	if len(path) == 0 {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	var inputID string

	slashIndex := strings.IndexRune(path, '/')
	if slashIndex == -1 {
		inputID = path
		if inputID == "all" {
			ViewAllToDo(res, req)
		} else {
			ViewToDo(inputID, res, req)
		}
	} else {
		res.WriteHeader(http.StatusNotFound)
	}
	return
}

//ViewAllToDo is handler function for Router /view/all
func ViewAllToDo(res http.ResponseWriter, req *http.Request) {
	writer := common.Writer{
		Resp: res,
	}

	data, err := common.GetAllToDo()
	if err != nil {
		writer.ResponseError(common.ErrorInternal)
	}

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})

}

//ViewToDo is handler function for Router /view/all
func ViewToDo(indexStr string, res http.ResponseWriter, req *http.Request) {

	writer := common.Writer{
		Resp: res,
	}

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		writer.ResponseError(common.ErrorIDInvalid)
		return
	}

	data, err := common.GetToDo(index)
	if err != nil {
		writer.ResponseError(common.ErrorTodoNotFound)
		return
	}

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    data,
	})

}
