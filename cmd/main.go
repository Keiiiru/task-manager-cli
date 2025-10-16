package main

import (
	"fmt"

	"github.com/task-manager-cli/internal/tasks"
)

func main() {
	repo := tasks.NewTaskRepository()
	service := tasks.NewTaskService(repo)
	handler := tasks.NewTaskHandler(service)

	for {
		fmt.Print("> ")
		var input string
		fmt.Scan(&input)

		switch input {
		case "add":
			{
				fmt.Println("Type task")
				fmt.Print("> ")
				var task string
				fmt.Scan(&task)
				handler.AddHandler(task)
			}
		case "remove":
			{
				fmt.Println("Type id")
				fmt.Print("> ")
				var id string
				fmt.Scan(&id)
				handler.RemoveHandler(id)
			}
		case "list":
			{
				handler.ListHandler()
			}
		}

		if input == "exit" {
			fmt.Println("Ending program...")
			break
		}
	}
}
