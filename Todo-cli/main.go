package main

import (
	"fmt"
)

//	func printTodos(todos Todos) {
//		for i, t := range todos {
//			completedAt := "nil"
//			if t.completedAt != nil {
//				completedAt = t.completedAt.Format("2006-01-02 15:04:05")
//			}
//			fmt.Printf("%d. %s | Completed: %v | CreatedAt: %s | CompletedAt: %s\n",
//				i, t.Title, t.Completed, t.createdAt.Format("2006-01-02 15:04:05"), completedAt)
//		}
//		fmt.Println()
//	}
func main() {
	fmt.Println("Inside main in Todo App")

	todos := Todos{}
	// todos.add("Todo 1")
	// todos.add("Todo 2")
	// todos.add("Todo 3")

	// todos.toggle(0)
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	cmdFlage := NewCmdFlags()
	cmdFlage.Execute(&todos)

	storage.Save(todos)

	// todos.print()
}
