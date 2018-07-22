package api

import (
	"net/http"
	bl "github.com/rathodc/todo/task/bl"
)

func GetTaskListHandler(w http.ResponseWriter, r *http.Request)  {
    w.WriteHeader(http.StatusOK)
    res:= bl.GetTaskList(r.Header.Get("Auth-Token"))
    w.Header().Set("Content-Type", "application/json")
    byt:=[]byte(res)
    w.Write(byt)
}
