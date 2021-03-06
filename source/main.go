package main

import (
	"log"
	"net/http"

	"three-sixty-to-do/source/common"
	"three-sixty-to-do/source/handler"
)

func init() {
	common.TodoMap = make(map[int]common.TodoDetail)
	common.UID = 1
}

func main() {
	log.Println("[*]Service started")
	http.HandleFunc("/add", handler.AddToDo)
	http.HandleFunc("/view/", handler.ViewHandler)
	http.HandleFunc("/remove", handler.RemoveToDo)
	http.HandleFunc("/modify", handler.ModifyToDo)
	http.HandleFunc("/modify/status", handler.ModifyStatusToDo)

	http.ListenAndServe(":1112", nil)
}
