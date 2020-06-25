package main

func createID(c ClientResponse) int {
	// Create newID = lasted Todos ID + 1
	if index := len(c.Todos); index > 0 {
		return c.Todos[index].ID + 1 
	}
	return 0
}

func (c ClientResponse) deleteTodo(id int)  {
	for i, t := range c.Todos {
		if t.ID == id {
			c.Todos = append(c.Todos[:i], c.Todos[i+1:0]...)
		}
	}	
}

func (c ClientResponse) clearTodo() {
	j := 0
			for _, t := range c.Todos {
				if (t.Completed) {
					c.Todos[j] = t	
					j++
				}
			}
			c.Todos = c.Todos[:j]
}

func (c ClientResponse) toggleCompleteTodo(id int)  {
	for i, t := range c.Todos {
		if t.ID == id {
			c.Todos[i].Completed = !c.Todos[i].Completed
		}
	}	
}

func (c ClientResponse) toggleCompleteTask(loadID [2]int)  {
	// Loop through Todos list -> find selected Todo -> loop through selected todo 
	// -> find selected task -> toggle selected task status
	for i, todo := range c.Todos {
		if todo.ID == loadID[0] {
		 for j, task := range c.Todos[i].Tasks {
				if task.TaskID == loadID[1] {
					c.Todos[i].Tasks[j].IsDone = !c.Todos[i].Tasks[j].IsDone	
				}	
			}	
		}
	}	
}

func (c ClientResponse) deletetask(loadID [2]int)  {
	// Loop through Todos list -> find selected Todo -> loop through selected todo 
	// -> find selected task -> delete selected task 
	for i, todo := range c.Todos {
		if todo.ID == loadID[0] {
			for j, task := range c.Todos[i].Tasks {
				if task.TaskID == loadID[1] {
					c.Todos[i].Tasks = append(c.Todos[i].Tasks[:j], c.Todos[i].Tasks[j+1:]...) 
				}
			}
		}
	}
}