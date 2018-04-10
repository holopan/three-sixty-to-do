package main

import (
	"log"
	"net/http"

	"to-do-list/source/common"
	"to-do-list/source/handler"
)

func init() {

}

func main() {
	log.Println("[*]Service started")
	http.HandleFunc("/", welcome)
	http.HandleFunc("/add", handler.AddToDo)
	http.HandleFunc("/view/", handler.ViewHandler)

	http.ListenAndServe(":1112", nil)
}

func welcome(res http.ResponseWriter, req *http.Request) {
	writer := common.Writer{
		Resp: res,
	}

	writer.Response(common.SuccessResponse{
		Status:  true,
		Message: "Welcome",
	})

}
