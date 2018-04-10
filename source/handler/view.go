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
		return
	} else {
		res.WriteHeader(http.StatusNotFound)
		return
	}
}

//ViewAllToDo is handler function for Router /view/all
func ViewAllToDo(res http.ResponseWriter, req *http.Request) {
	writer := common.Writer{
		Resp: res,
	}

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    common.TodoList,
	})

}

//ViewToDo is handler function for Router /view/all
func ViewToDo(indexStr string, res http.ResponseWriter, req *http.Request) {
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		//return error
		return
	}

	if common.TodoList == nil {
		//return error
		return
	} else if index < 0 {
		//return error
		return
	} else if len(common.TodoList) < index {
		//return error
		return
	}
	writer := common.Writer{
		Resp: res,
	}

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    common.TodoList[index],
	})

}
