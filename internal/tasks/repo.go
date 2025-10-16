package tasks

import "log"

type Repository interface {
	AddTask(task TaskDTO) (string, error)
	RemoveTask(id int) (string, error)
	ListTasks() (map[int]TaskDTO, error)
}

type TaskRepository struct {
	storage   map[int]TaskDTO
	currentId int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{storage: map[int]TaskDTO{}}
}

func (r *TaskRepository) AddTask(task TaskDTO) (string, error) {
	log.Printf("[INFO]: Set task into storage\n")

	r.currentId++
	task.Id = r.currentId
	r.storage[r.currentId] = task

	return "Successful", nil
}

func (r *TaskRepository) RemoveTask(id int) (string, error) {
	log.Printf("[INFO]: Deleting task from storage\n")
	delete(r.storage, id)

	return "Successful", nil
}

func (r *TaskRepository) ListTasks() map[int]TaskDTO {
	return r.storage
}
