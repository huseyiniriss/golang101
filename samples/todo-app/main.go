package main

import (
	"fmt"
)

type ToDo struct {
	Id    int
	Title string
	Done  bool
}

var todoList []ToDo

func main() {
	// addTodo
	addTodo(ToDo{Title: "Learn Go", Done: false})
	addTodo(ToDo{Title: "Learn Java", Done: false})
	addTodo(ToDo{Title: "Learn Python", Done: false})
	addTodo(ToDo{Title: "Learn C++", Done: false})
	addTodo(ToDo{Title: "Learn C", Done: false})

	for {
		fmt.Println("--- TODO APP ---")
		fmt.Println("1. Add Todo")
		fmt.Println("2. Remove Todo")
		fmt.Println("3. Update Todo")
		fmt.Println("4. View Todo")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		condition := 0
		fmt.Scanf("%d", &condition)
		switch condition {
		case 1:
			todo := ToDo{}
			fmt.Print("Enter title: ")
			fmt.Scanf("%s", &todo.Title)
			fmt.Print("Enter done: ")
			fmt.Scanf("%t", &todo.Done)
			addTodo(todo)
		case 2:
			var i int
			fmt.Print("Enter id: ")
			fmt.Scanf("%d", &i)
			removeTodo(i)
		case 3:
			var i int
			fmt.Print("Enter id: ")
			fmt.Scanf("%d", &i)
			todo := ToDo{}
			fmt.Print("Enter title: ")
			fmt.Scanf("%s", &todo.Title)
			fmt.Print("Enter done: ")
			fmt.Scanf("%t", &todo.Done)
			updateTodo(i, todo)
		case 4:
			viewTodo()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addTodo(todo ToDo) {
	if todo.Title == "" {
		fmt.Println("Invalid title")
		return
	}
	if todo.Done != true && todo.Done != false {
		fmt.Println("Invalid done")
		return
	}
	todo.Id = len(todoList)

	todoList = append(todoList, todo)
}

func removeTodo(id int) {
	for i, todo := range todoList {
		if todo.Id == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
			return
		}
	}
	fmt.Println("Invalid id")
}

func updateTodo(id int, todo ToDo) {
	for i, item := range todoList {
		if item.Id == id {
			todo.Id = id
			todoList[i] = todo
			return
		}
	}
	fmt.Println("Invalid id")
}

func viewTodo() {
	fmt.Println("Id\tTitle\tDone")
	for i, todo := range todoList {
		fmt.Println(i, todo)
	}
}
