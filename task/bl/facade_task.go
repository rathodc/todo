package bl

import (
	"math/rand"
	"encoding/hex"
	"encoding/json"
	dal "github.com/rathodc/todo/task/dal"
)


func AddToTaskList(user string,tsk_str string){
	var task_obj dal.Task
	byt:= []byte(tsk_str)
	json.Unmarshal(byt, &task_obj)
	b := make([]byte, 4) //equals 8 charachters
	rand.Read(b)
	tid:= hex.EncodeToString(b)
	task_obj.Tid = tid
	dal.AddToTaskList(user,task_obj)
}

func RemoveTask(user string,tid string) bool {
	return dal.RemoveTask(user,tid)
}

func RemoveTaskList(user string) bool {
	return dal.RemoveTaskList(user)
}


func GetTaskList(user string) string {
	tsk_list:= dal.GetTaskList(user)
	str,err:=json.Marshal(tsk_list)
	if err!=nil{
		return "[]"
	} else {
		return string(str)
	}
}

func GetTask(user string,tid string) (string,bool) {
	var task_obj = dal.GetTask(user,tid)
	if task_obj.Tid == "" {
		return "{}",false
	} else {
		str,err:=json.Marshal(task_obj)
		if err != nil {
			return "{}",false
		} else {
			return string(str),true
		}
	}
}

func UpdateTask(user string,tid string,tsk_str string) (string,bool) {
	var task_obj dal.Task
	byt:= []byte(tsk_str)
	json.Unmarshal(byt, &task_obj)
	updated_task_obj:=dal.UpdateTask(user,tid,task_obj)

	if updated_task_obj.Tid == "" {
		return "{}",false
	} else {
		str,err:=json.Marshal(updated_task_obj)
		if err == nil {
			return string(str),true
		} else {

			return "{}",false
		}
	}
}



