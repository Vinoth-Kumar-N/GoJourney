package main

import (
	"fmt"
)


func main() {
	fmt.Println("Inside main in Todo App")

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	cmdFlage := NewCmdFlags()
	cmdFlage.Execute(&todos)

	storage.Save(todos)

}






























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