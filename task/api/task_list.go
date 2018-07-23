package api

import (
	"net/http"
	bl "github.com/rathodc/todo/task/bl"
)

func GetTaskListHandler(w http.ResponseWriter, r *http.Request)  {
    res:= bl.GetTaskList(r.Header.Get("Auth-Token"))
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    byt:=[]byte(res)
    w.Write(byt)
}

func RemoveTaskListHandler(w http.ResponseWriter, r *http.Request)  {
    bl.RemoveTaskList(r.Header.Get("Auth-Token"))
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNoContent)
}