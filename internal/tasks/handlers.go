package tasks

import (
	"fmt"
	"strconv"
)

type HandlerInterface interface {
	AddHandler(task string) (string, error)
}

type TaskHandler struct {
	service *TaskService
}

func NewTaskHandler(s *TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) AddHandler(task string) (string, error) {
	fmt.Printf("[INFO]: handler func\n")

	h.service.AddService(task)

	return "Success", nil
}

func (h *TaskHandler) RemoveHandler(id string) (string, error) {
	fmt.Printf("[INFO]: remove handler")
	intId, err := strconv.Atoi(id)

	if err != nil {
		return "", fmt.Errorf("[ERROR]: failed convert string to int")
	}

	h.service.RemoveTask(intId)
	return "success", nil
}

func (h TaskHandler) ListHandler() (map[int]string, error) {
	tasks, _ := h.service.ListTasks()

	return tasks, nil
}
