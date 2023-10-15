package store

import (
	"errors"

	"github.com/Atsu77/go_todo_app/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[entity.TaskId]*entity.Task{}}
	ErrNotFound = errors.New("task not found")
)

type TaskStore struct {
	LastId entity.TaskId
	Tasks map[entity.TaskId]*entity.Task
}

func (ts * TaskStore) Add(t *entity.Task) (int, error) {
	ts.LastId++
	t.ID = ts.LastId
	ts.Tasks[t.ID] = t
	return int(t.ID), nil
}

func (ts * TaskStore) All() entity.Tasks {
	tasks := make(entity.Tasks, 0, len(ts.Tasks))
	for _, t := range ts.Tasks {
		tasks = append(tasks, t)
	}
	return tasks
}