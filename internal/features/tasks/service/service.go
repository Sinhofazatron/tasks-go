package tasks_service

import (
	"context"

	"github.com/Sinhofazatron/tasks-go/internal/core/domain"
)

type TasksService struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTask(ctx context.Context, id int) (domain.Task, error)
	PatchTask(ctx context.Context, id int, patch domain.Task) (domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
	GetTasks(ctx context.Context, userID, limit, offset *int) ([]domain.Task, error)
}

func NewTasksService(tasksRepository TasksRepository) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
