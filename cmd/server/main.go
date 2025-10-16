package main

import (
	"log"
	"net/http"

	"github.com/task-manager/internal/tasks"
)

func main() {
	repo := tasks.NewTaskRepository()
	service := tasks.NewTaskService(repo)
	handler := tasks.NewTaskHandler(service)

	mux := http.NewServeMux()

	mux.HandleFunc("/tasks/add", handler.AddHandler)
	mux.HandleFunc("/tasks/{id}", handler.RemoveHandler)
	mux.HandleFunc("/tasks", handler.ListHandler)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
