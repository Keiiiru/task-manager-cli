package tasks

import "fmt"

type Repository interface {
	AddTask(task string) (string, error)
	RemoveTask(id int) (string, error)
	ListTasks() (map[int]string, error)
}

type TaskRepository struct {
	storage   map[int]string
	currentId int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{storage: map[int]string{}, currentId: 0}
}

func (r *TaskRepository) AddTask(task string) (string, error) {
	fmt.Printf("[INFO]: repository add task\n")

	r.currentId++
	r.storage[r.currentId] = task
	fmt.Println(r.storage)

	return "Successful", nil
}

func (r *TaskRepository) RemoveTask(id int) (string, error) {
	fmt.Printf("[INFO]: deleting task\n")
	delete(r.storage, id)

	return "Successful", nil
}

func (r TaskRepository) ListTasks() (map[int]string, error) {
	return r.storage, nil
}
