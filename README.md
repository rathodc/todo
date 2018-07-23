# todo application

### This is a simple To DO application with following features :-<br>
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

### 6.Launch Application
```
chirag@chirag-rathod:~/test_assignment/src/github.com/rathodc/todo/app$ app
Application will run at port 8000
```

# API Usage

**Create/Add a Task**
----
  This api will help you add a task in your "to-do" list

* **URL**

  /api/v1/todo/task

* **Method:**
  
  `POST`
  
*  **Header**

   All api calls requires an authorization which uniquely identifies a particular user.For this application use one of these auth tokens - "e4b76ae67e4b","604605a6bfe0","22437a5dc2f5"

   **Required:**
 
   `Auth-Token=[Token]`<br>
   `Content-Type=application/json`

* **Body**

  E.g json payload - {"tname":"Create Design","end_date":"2019-09-12","notify":60}
  

* **Success Response:**
  
  {"success":true}

  * **Code:** 200 <br />
 
* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{"error_title":"Authorization_Unavailable","error_message":"Please provide Valid Auth Token"}`



**Get Task List**
----
  This api will help you fetch all tasks from your todo application

* **URL**

  /api/v1/todo/task_list

* **Method:**
  
  `GET`
  
*  **Header**

   All api calls requires an authorization which uniquely identifies a particular user.For this application use one of these auth tokens - "e4b76ae67e4b","604605a6bfe0","22437a5dc2f5"

   **Required:**
 
   `Auth-Token=[Token]`


* **Success Response:**
  
  [
    {
        "tid": "52fdfc07",
        "tname": "3sdqdqdlay f;nwnf",
        "end_date": "2019-09-12",
        "notify": 60
    }
  ]

  * **Code:** 200 <br />
 
* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{"error_title":"Authorization_Unavailable","error_message":"Please provide Valid Auth Token"}`

**Update an existing task**
----
  This api will help you update a task in your "to-do" list

* **URL**

  /api/v1/todo/task/{tid}

* **Method:**
  
  `PUT`
  
*  **Header**

   All api calls requires an authorization which uniquely identifies a particular user.For this application use one of these auth tokens - "e4b76ae67e4b","604605a6bfe0","22437a5dc2f5"

   **Required:**
 
   `Auth-Token=[Token]`<br>
   `Content-Type=application/json`

* **Body**

  E.g json payload - {"tname": "Implementation"}
  

* **Success Response:**
  
  {"tid":"52fdfc07","tname":"Implementation","end_date":"2019-09-12","notify":60}

  * **Code:** 200 <br />
 
* **Error Response:**

  * **Code:** 404 Not Found <br />
    **Content:** `{"error_title":"Bad_Request","error_message":"Task not found"}

**Get information about an existing task**
----
  This api will help you get information about a particular task in your "to-do" list

* **URL**

  /api/v1/todo/task/{tid}

* **Method:**
  
  `GET`
  
*  **Header**

   All api calls requires an authorization which uniquely identifies a particular user.For this application use one of these auth tokens - "e4b76ae67e4b","604605a6bfe0","22437a5dc2f5"

   **Required:**
 
   `Auth-Token=[Token]`<br>


* **Success Response:**
  
  {"tid":"52fdfc07","tname":"Implementation","end_date":"2019-09-12","notify":60}

  * **Code:** 200 <br />
 
* **Error Response:**

  * **Code:** 404 Not Found <br />
    **Content:** `{"error_title":"Bad_Request","error_message":"Task not found"}

**Delete an existing task**
----
  This api will help you delete an existing task from your "to-do" list

* **URL**

  /api/v1/todo/task/{tid}

* **Method:**
  
  `DELETE`
  
*  **Header**

   All api calls requires an authorization which uniquely identifies a particular user.For this application use one of these auth tokens - "e4b76ae67e4b","604605a6bfe0","22437a5dc2f5"

   **Required:**
 
   `Auth-Token=[Token]`

* **Success Response:**
  
  No response content

  * **Code:** 204 <br />
 
* **Error Response:**

  * **Code:** 404 Not Found <br />
    **Content:** `{"error_title":"Bad_Request","error_message":"Task not found"}

**Delete the task list**
----
  This api will help you delete all tasks from your list

* **URL**

  /api/v1/todo/task_list

* **Method:**
  
  `DELETE`
  
*  **Header**

   All api calls requires an authorization which uniquely identifies a particular user.For this application use one of these auth tokens - "e4b76ae67e4b","604605a6bfe0","22437a5dc2f5"

   **Required:**
 
   `Auth-Token=[Token]`

* **Success Response:**
  
  No response content

  * **Code:** 204 <br />
 
* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `{"error_title":"Authorization_Unavailable","error_message":"Please provide Valid Auth Token"}`

# Unit Tests

### Go to app directory to run your unit tests
```
chirag@chirag-rathod:~/test_assignment/src/github.com/rathodc/todo/app$ go test
```

