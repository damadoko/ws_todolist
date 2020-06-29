package main

import (
	"math"
)

func createID(c *ClientResponse) int {
	// Create newID = lasted Todos ID + 1
	if index := len(c.Todos); index > 0 {
		return c.Todos[index-1].ID + 1 
	}
	return 0
}

func (c *ClientResponse) deleteTodo(id int)  {
	for i, t := range c.Todos {
		if t.ID == id {
			c.Todos = append(c.Todos[:i], c.Todos[i+1:]...)
		}
	}	
}

func (c *ClientResponse) clearTodo() {
	j := 0
			for _, t := range c.Todos {
				if (!t.Completed) {
					c.Todos[j] = t	
					j++
				}
			}
			c.Todos = c.Todos[:j]
}

func (c *ClientResponse) toggleCompleteTodo(id int)  {
	for i, t := range c.Todos {
		if t.ID == id {
			c.Todos[i].Completed = !c.Todos[i].Completed
			// Change all Task's isDone to true or false
				for _, task := range c.Todos[i].Tasks {
					task.IsDone = (c.Todos[i].Completed)
				}
			// Update percentage
			c.Todos[i].updatePercentage()
		}
	}	
}

func (c *ClientResponse) toggleCompleteTask(loadID [2]int)  {
	// Loop through Todos list -> find selected Todo -> loop through selected todo 
	// -> find selected task -> toggle selected task status
	for i, todo := range c.Todos {
		if todo.ID == loadID[0] {
		 for j, task := range c.Todos[i].Tasks {
				if task.TaskID == loadID[1] {
					c.Todos[i].Tasks[j].IsDone = !c.Todos[i].Tasks[j].IsDone	
				}	
			}	
			// Update percentage
			c.Todos[i].updatePercentage()
			// Update complement
			c.Todos[i].updateComplete()
		}
	}	
}

func (c *ClientResponse) deletetask(loadID [2]int)  {
	// Loop through Todos list -> find selected Todo -> loop through selected todo 
	// -> find selected task -> delete selected task 
	for i, todo := range c.Todos {
		if todo.ID == loadID[0] {
			for j, task := range c.Todos[i].Tasks {
				if task.TaskID == loadID[1] {
					c.Todos[i].Tasks = append(c.Todos[i].Tasks[:j], c.Todos[i].Tasks[j+1:]...) 
				}
			}
			// Update percentage
			c.Todos[i].updatePercentage()
			// Update complement
			c.Todos[i].updateComplete()
		}
	}
}

func (t *Todo) updatePercentage () {
	if t.Completed {
		t.Percentage = 100
	} else {
		doneSlice := []Task{}
		for _,task := range t.Tasks {
			if task.IsDone {
				doneSlice = append(doneSlice, task)
			}
		}
		rawPercentage := (float64(len(doneSlice)) /float64(len(t.Tasks)) )
		t.Percentage = math.Round(rawPercentage * 10000) /100
	}
}

func (t *Todo) updateComplete()  {
	if t.Percentage == 100 {
		t.Completed = true
	} else {
		t.Completed = false
	}
}