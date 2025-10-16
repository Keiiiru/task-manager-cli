package tasks

import "fmt"

type Service interface {
	AddService(task string) (string, error)
}

type TaskService struct {
	repo *TaskRepository
}

func NewTaskService(r *TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) AddService(task string) (string, error) {
	fmt.Printf("[INFO]: add service\n")

	if _, err := s.repo.AddTask(task); err != nil {
		return "", fmt.Errorf("[ERROR]: something went wrong")
	}

	return "Success", nil
}

func (s *TaskService) RemoveTask(id int) (string, error) {
	fmt.Printf("[INFO]: deleting task service\n")

	s.repo.RemoveTask(id)

	return "Successful", nil
}

func (s TaskService) ListTasks() (map[int]string, error) {
	fmt.Printf("[INFO]: getting tasks storage\n")

	tasks, _ := s.repo.ListTasks()

	return tasks, nil
}
