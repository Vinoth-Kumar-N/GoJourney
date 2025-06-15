package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) delete(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completedTime := time.Now()
		t[index].CompletedAt = &completedTime
	}

	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	t[index].CompletedAt = nil
	t[index].Completed = false
	return nil
}
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of range")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *todos {
		completed := "❌"
		CompletedAt := ""
		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				CompletedAt = t.CompletedAt.Format("2006-01-02 15:04:05")
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format("2006-01-02 15:04:05"), CompletedAt)
	}
	table.Render()
}
