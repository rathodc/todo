package api

import (
	"net/http"
	"log"
	"io/ioutil"
	"os"
	"github.com/gorilla/mux"
	"github.com/xeipuuv/gojsonschema"
	bl "github.com/rathodc/todo/task/bl"
)

func validate_req_body(req_body string,schema_file string) bool {
	schemaLoader := gojsonschema.NewReferenceLoader(schema_file)
	documentLoader := gojsonschema.NewStringLoader(req_body)
        result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
            log.Println(err.Error())
        }
	return result.Valid()
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request)  {
    b, _ := ioutil.ReadAll(r.Body)
    dir, _ := os.Getwd()
    schema_file:="file:///"+dir+"/schema_create.json"

    if validate_req_body(string(b),schema_file) {
	    bl.AddToTaskList(r.Header.Get("Auth-Token"), string(b))
	    w.WriteHeader(http.StatusOK)
	    w.Write([]byte(`{"success":true}`))
    } else {
	    http.Error(w, `{"error_title":"Bad_Request","error_message":"Invalid Request"}`, http.StatusBadRequest)
    }
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request)  {
	b, _ := ioutil.ReadAll(r.Body)
	dir, _ := os.Getwd()
    	schema_file:="file:///"+dir+"/schema_update.json"

	if validate_req_body(string(b),schema_file) {
		params := mux.Vars(r)
		res,status:= bl.UpdateTask(r.Header.Get("Auth-Token"), params["tid"], string(b))
		w.Header().Set("Content-Type", "application/json")
		byt := []byte(res)
		if !status {
			http.Error(w, `{"error_title":"Bad_Request","error_message":"Task not found"}`, http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(byt)
		}

	} else {
	    http.Error(w, `{"error_title":"Bad_Request","error_message":"Invalid Request"}`,http.StatusBadRequest)
    	}
}


func GetTaskHandler(w http.ResponseWriter, r *http.Request)  {
	params:= mux.Vars(r)
	res,status:=bl.GetTask(r.Header.Get("Auth-Token"),params["tid"])
	w.Header().Set("Content-Type", "application/json")
	if !status {
		http.Error(w, `{"error_title":"Bad_Request","error_message":"Task not found"}`, http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	}
}


func RemoveTaskHandler(w http.ResponseWriter, r *http.Request)  {
	params:= mux.Vars(r)
	res:=bl.RemoveTask(r.Header.Get("Auth-Token"),params["tid"])
	w.Header().Set("Content-Type", "application/json")
	if res {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, `{"error_title":"Bad_Request","error_message":"Task not found"}`, http.StatusNotFound)
	}
}
