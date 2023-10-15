package handler

import (
	"net/http"

	"github.com/Atsu77/go_todo_app/entity"
	"github.com/Atsu77/go_todo_app/store"
)

type ListTask struct {
	Store *store.TaskStore
}

type task struct {
	ID     entity.TaskId     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks := lt.Store.All()
	res := make([]task, 0, len(tasks))
	for _, t := range tasks {
		res = append(res, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx, w, res, http.StatusOK)
}
