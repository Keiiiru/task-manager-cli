package tasks

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type HandlerInterface interface {
	AddHandler(w http.ResponseWriter, r *http.Request)
	RemoveHandler(w http.ResponseWriter, r *http.Request)
}

type TaskHandler struct {
	s *TaskService
}

func NewTaskHandler(service *TaskService) *TaskHandler {
	return &TaskHandler{service}
}

func (h *TaskHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	log.Println("[INFO]: Handle add task request")
	var task TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Fatalln("[ERROR]: invalid request body")

	}

	log.Println("[INFO]: request body successfuly parsed")

	defer r.Body.Close()

	h.s.AddService(task)

	log.Println("[INFO]: handler ends work")

}

func (h *TaskHandler) RemoveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	log.Println("[INFO]: Handle remove task request")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	h.s.RemoveTask(id)

	log.Println("[INFO]: handler ends work")
}

func (h *TaskHandler) ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	log.Println("[INFO]: Handle get list of tasks")

	tasks := h.s.ListTasks()
	json.NewEncoder(w).Encode(tasks)

	log.Println("[INFO]: handle ends work")
}
