# todo application

### This application supports a simple To DO application with following features :-<br>
1.Provides an endpoint to get a list of tasks<br>
2.Provides an endpoint to add a task to your "To-Do" list<br>
3.Provides an endpoint to modify a particular task in your "To-Do" list<br>
4.Provides an endpoint to delete a particular task from your "To-Do" list<br>
5.Get information of a particular task in your "To-Do" list<br>
<br>
### Each Task added to your "To-Do" list contains following information :-<br>
1.Task ID - This is generated automatically by the application<br>
2.Task Name - This is the name which user can select while creating a task<br>
3.End Date - This is the task completion date<br>
4.Notify - This is the duration(in mins) which user can set to notify him before the completion date is passed<br>


# Set-Up Instructions

Make sure you have go environment set up on your machine before proceeding to below steps<br>


### 1.Create a Workspace,say - test_assignment
```
chirag@chirag-rathod:~$ mkdir test_assignment
```

### 2.Export GOPATH
```
chirag@chirag-rathod:~$ export GOPATH=/home/chirag/test_assignment/
chirag@chirag-rathod:~$ cd test_assignment/
```

### 3.Get GIT repository
```
chirag@chirag-rathod:~/test_assignment$ go get github.com/rathodc/todo
package github.com/rathodc/todo: no Go files in /home/chirag/test_assignment/src/github.com/rathodc/todo
```

### 4.Get Dependencies
```
chirag@chirag-rathod:~/test_assignment$ go get github.com/gorilla/mux
chirag@chirag-rathod:~/test_assignment$ go get github.com/xeipuuv/gojsonschema
chirag@chirag-rathod:~/test_assignment$ go get github.com/stretchr/testify/assert
```

### 5.Install Application
```
chirag@chirag-rathod:~/test_assignment/src/github.com/rathodc/todo/app$ go install
chirag@chirag-rathod:~/test_assignment/src/github.com/rathodc/todo/app$ export PATH=$GOPATH/bin:$PATH
chirag@chirag-rathod:~/test_assignment/src/github.com/rathodc/todo/app$
```
