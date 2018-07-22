package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "log"
    "bytes"
    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    "io/ioutil"
    "encoding/json"
    api "github.com/rathodc/todo/task/api"
    dal "github.com/rathodc/todo/task/dal"
)

type task struct{
	Tid string `json:"tid"`
	Tname string `json:"tname"`
	End_Date string `json:"end_date"`
	Notify int `json:"notify"`
}

var tasks []task

func loggingNValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)
        token:=r.Header.Get("Auth-Token")
        if !dal.IsValidToken(token) {
            http.Error(w, `{"error_title":"Authorization_Unavailable","error_message":"Please provide Valid Auth Token"}`, http.StatusForbidden)
        } else {
            next.ServeHTTP(w, r)
        }
    })
}

func Router() *mux.Router {
    router := mux.NewRouter()
    router.Use(loggingNValidationMiddleware)
    router.HandleFunc("/api/v1/todo/task_list", api.GetTaskListHandler).Methods("GET")
    router.HandleFunc("/api/v1/todo/task", api.AddTaskHandler).Methods("POST")
    router.HandleFunc("/api/v1/todo/task/{tid}", api.UpdateTaskHandler).Methods("PUT")
    router.HandleFunc("/api/v1/todo/task/{tid}", api.RemoveTaskHandler).Methods("DELETE")
    router.HandleFunc("/api/v1/todo/task/{tid}", api.GetTaskHandler).Methods("GET")
    return router
}

func TestGetTaskList(t *testing.T) {
    dal.InitializeUsers()
    request, _ := http.NewRequest("GET", "/api/v1/todo/task_list", nil)
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestCreateTaskPass(t *testing.T) {
    dal.InitializeUsers()


    request, _ := http.NewRequest("GET", "/api/v1/todo/task_list", nil)
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    res_body,_:=ioutil.ReadAll(response.Body)
    json.Unmarshal(res_body,&tasks)
    assert.Equal(t, 0, len(tasks), "Initially Length should be 0")

    req_body:= []byte(`{"tname":"Create Design","end_date":"2018-09-01","notify":60}`)
    request, _ = http.NewRequest("POST", "/api/v1/todo/task", bytes.NewBuffer(req_body))
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response = httptest.NewRecorder()
    Router().ServeHTTP(response, request)

    request, _ = http.NewRequest("GET", "/api/v1/todo/task_list", nil)
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response = httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    res_body,_=ioutil.ReadAll(response.Body)
    json.Unmarshal(res_body,&tasks)
    assert.Equal(t, 200, response.Code, "OK response is expected")
    assert.Equal(t, 1, len(tasks), "Now the length should be 1")
}


func TestCreateTaskFail(t *testing.T) {
    dal.InitializeUsers()
    var r_body = []byte(`{"tname":"Create Design","end_date":"2018-09-01"}`)
    request, _ := http.NewRequest("POST", "/api/v1/todo/task", bytes.NewBuffer(r_body))
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 400, response.Code, "Bad Request expected")
}


func TestUpdateTask(t *testing.T) {
    dal.InitializeUsers()


    req_body:= []byte(`{"tname":"Create Design","end_date":"2018-09-01","notify":60}`)
    request, _ := http.NewRequest("POST", "/api/v1/todo/task", bytes.NewBuffer(req_body))
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response:= httptest.NewRecorder()
    Router().ServeHTTP(response, request)

    request, _ = http.NewRequest("GET", "/api/v1/todo/task_list", nil)
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response = httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    res_body,_:=ioutil.ReadAll(response.Body)
    json.Unmarshal(res_body,&tasks)

    new_tname:= "Implement"
    tid:=tasks[0].Tid
    tasks[0].Tname = new_tname

    json_str,_:=json.Marshal(tasks[0])
    request, _ = http.NewRequest("PUT", "/api/v1/todo/task/"+tid, bytes.NewBuffer(json_str))
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response= httptest.NewRecorder()
    Router().ServeHTTP(response, request)

    res_body,_=ioutil.ReadAll(response.Body)
    json.Unmarshal(res_body,&tasks)

    assert.Equal(t, new_tname, tasks[0].Tname, "Updated task name is expected")
}

func TestDeleteTask(t *testing.T) {
    dal.InitializeUsers()

    req_body:= []byte(`{"tname":"Create Design","end_date":"2018-09-01","notify":60}`)
    request, _ := http.NewRequest("POST", "/api/v1/todo/task", bytes.NewBuffer(req_body))
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response:= httptest.NewRecorder()
    Router().ServeHTTP(response, request)

    request, _ = http.NewRequest("GET", "/api/v1/todo/task_list", nil)
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response = httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    res_body,_:=ioutil.ReadAll(response.Body)
    json.Unmarshal(res_body,&tasks)

    tid:=tasks[0].Tid
    assert.Equal(t, 1, len(tasks), "Now the length should be 1")

    request, _ = http.NewRequest("DELETE", "/api/v1/todo/task/"+tid, nil)
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response= httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 204, response.Code, "No response status code is expected")

    request, _ = http.NewRequest("GET", "/api/v1/todo/task_list", nil)
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response = httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    res_body,_=ioutil.ReadAll(response.Body)
    json.Unmarshal(res_body,&tasks)
    assert.Equal(t, 0, len(tasks), "Total tasks should be 0")
    assert.Equal(t, 200, response.Code, "OK response is expected")
}


func TestGetTask(t *testing.T) {
    dal.InitializeUsers()

    req_body := []byte(`{"tname":"Create Design","end_date":"2018-09-01","notify":60}`)
    request, _ := http.NewRequest("POST", "/api/v1/todo/task", bytes.NewBuffer(req_body))
    request.Header.Set("Auth-Token", "e4b76ae67e4b")
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)

    request, _ = http.NewRequest("GET", "/api/v1/todo/task_list", nil)
    request.Header.Set("Auth-Token", "e4b76ae67e4b")
    response = httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    res_body, _ := ioutil.ReadAll(response.Body)
    json.Unmarshal(res_body, &tasks)

    tid:=tasks[0].Tid
    request, _ = http.NewRequest("GET", "/api/v1/todo/task/"+tid, nil)
    request.Header.Set("Auth-Token","e4b76ae67e4b")
    response= httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    res_body, _ = ioutil.ReadAll(response.Body)
    json.Unmarshal(res_body, &tasks)

    tid_from_api_output:= tasks[0].Tid

    assert.Equal(t, 200, response.Code, "OK response is expected")
    assert.Equal(t, tid, tid_from_api_output, "Task ID that has been passed to the api should match with output")

}