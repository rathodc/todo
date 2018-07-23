package libservice

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	api "github.com/rathodc/todo/task/api"
	dal "github.com/rathodc/todo/task/dal"
)

func loggingNValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)
        token:=r.Header.Get("Auth-Token")
        log.Println(token)
        if !dal.IsValidToken(token) {
            http.Error(w, `{"error_title":"Authorization_Unavailable","error_message":"Please provide Valid Auth Token"}`, http.StatusUnauthorized)
        } else {
            next.ServeHTTP(w, r)
        }
    })
}

func CreateApp()  {
	dal.InitializeUsers()
	router := mux.NewRouter()
	router.Use(loggingNValidationMiddleware)
	router.HandleFunc("/api/v1/todo/task_list", api.GetTaskListHandler).Methods("GET")
	router.HandleFunc("/api/v1/todo/task", api.AddTaskHandler).Methods("POST")
	router.HandleFunc("/api/v1/todo/task/{tid}", api.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/api/v1/todo/task/{tid}", api.RemoveTaskHandler).Methods("DELETE")
	router.HandleFunc("/api/v1/todo/task/{tid}", api.GetTaskHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}