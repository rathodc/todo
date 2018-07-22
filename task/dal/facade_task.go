package dal

type Task struct{
	Tid string `json:"tid"`
	Tname string `json:"tname"`
	End_Date string `json:"end_date"`
	Notify int `json:"notify"`
}


var task_list map[string][]Task

func InitializeUsers()  {
	task_list = make(map[string][]Task)
	task_list["e4b76ae67e4b"] = make([]Task,0)
	task_list["604605a6bfe0"] = make([]Task,0)
	task_list["22437a5dc2f5"] = make([]Task,0)
}

func IsValidToken(token string) bool  {
	_,found:= task_list[token]
	if found {
		return true
	}
	return false
}


func remove_index(s []Task, index int) []Task {
    return append(s[:index], s[index+1:]...)
}

func get_task_index(user string,tid string) int {
	for i:=range task_list[user] {
		if task_list[user][i].Tid == tid {
			return i
		}
	}
	return -1
}



func AddToTaskList(user string,task_obj Task)  {
	task_list[user] = append(task_list[user],task_obj)
}

func RemoveTask(user string,tid string) bool {
	for i:=range task_list[user] {
		if task_list[user][i].Tid == tid {
			task_list[user] = remove_index(task_list[user], i)
			return true
		}
	}
	return false
}

func GetTaskList(user string) []Task {
	task_obj,found:= task_list[user]
	if found {
		return task_obj
	} else {
		return make([]Task,0)
	}
}

func GetTask(user string,tid string) Task {
	for i:=range task_list[user] {
		if task_list[user][i].Tid == tid {
			return task_list[user][i]
		}
	}
	return Task{}
}

func UpdateTask(user string,tid string,task_obj Task) Task {
	task_index:=get_task_index(user,tid)
	if task_index!=-1 {
		if task_obj.Tname != "" {
			task_list[user][task_index].Tname = task_obj.Tname
		}

		if task_obj.End_Date != ""{
			task_list[user][task_index].End_Date = task_obj.End_Date
		}

		if task_obj.Notify !=0 {
			task_list[user][task_index].Notify = task_obj.Notify
		}

		return task_list[user][task_index]
	} else {
		return Task{}
	}
}


