package entity

type TaskId int64
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID      TaskId     `json:"id"`
	Title   string     `json:"title"`
	Status  TaskStatus `json:"status"`
	Created string     `json:"created"`
}

type Tasks []*Task