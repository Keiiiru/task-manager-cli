package tasks

import "log"

type Service interface {
	AddService(task TaskDTO)
	RemoveTask(id int)
	ListTasks() (map[int]TaskDTO, error)
}

type TaskService struct {
	repo *TaskRepository
}

func NewTaskService(r *TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) AddService(task TaskDTO) {
	log.Printf("[INFO]: AddService starts work\n")

	if _, err := s.repo.AddTask(task); err != nil {
		log.Fatalln("[ERROR]: something went wrong")
	}

	log.Println("[INFO]: AddService ends work")

}

func (s *TaskService) RemoveTask(id int) {
	log.Printf("[INFO]: Delete service starts work\n")

	s.repo.RemoveTask(id)

	log.Println("[INFO]: Delete service ends work")
}

func (s *TaskService) ListTasks() map[int]TaskDTO {
	log.Printf("[INFO]: getting tasks storage\n")

	tasks := s.repo.ListTasks()

	return tasks
}
