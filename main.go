package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type (
	// Task is needed step to comple Todo
	Task struct {
		IsDone    bool   `json:"isDone"`
		TaskTitle string `json:"taskTitle"`
		TaskID    int    `json:"taskID"`
	}
	// Todo is belong to a slice of Todos in mainState
	Todo struct {
		ID         int    `json:"id"`
		Title      string `json:"title"`
		Completed  bool   `json:"completed"`
		Percentage int    `json:"percentage"`
		Tasks      []Task `json:"tasks"`
	}
	// ClientRequest store request from client
	ClientRequest struct {
		Type   string `json:"type,omitempty"` // type: add, delete, clear, changeFilter, setComplete
		LoadID [2]int  `json:"loadID,omitempty"` //loadID = [1, 2] -> ID: 0, taskID: 2
		Filter string `json:"filter,omitempty"`
		Todo   Todo   `json:"todo,omitempty"`
	}
	// ClientResponse store response from this server to client
	ClientResponse struct {
		Filter string `json:"filter"`
		Todos  []Todo `json:"todos"`
	}
)
var upgrader websocket.Upgrader
var updatedState = ClientResponse{
	Filter: "all",
	Todos: []Todo{} ,
}

func main() {
	fmt.Println("Server running...")
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":8001", nil)) 
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	
	for {
		clientReq := &ClientRequest{}
		err := conn.ReadJSON(clientReq)
		if err != nil {
						log.Println(err)
						return
		}
		
		clientResp := &ClientResponse{}
		switch clientReq.Type {
		case "add":
			clientReq.Todo.ID = createID(updatedState)
			updatedState.Todos = append(updatedState.Todos, clientReq.Todo)	

		case "delete":
			updatedState.deleteTodo(clientReq.LoadID[0])
			
		case "clear":
			updatedState.clearTodo()

		case "changeFilter":
			updatedState.Filter = clientReq.Filter	

		case "completeTodo":
			updatedState.toggleCompleteTodo(clientReq.LoadID[0])

		case "completeTask":
			updatedState.toggleCompleteTask(clientReq.LoadID)
			
		case "delTask":
			updatedState.deletetask(clientReq.LoadID)

			*clientResp = updatedState
		}
		if err := conn.WriteJSON(clientResp); err != nil {
						log.Println(err)
						return
		}
}
}