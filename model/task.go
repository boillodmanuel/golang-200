package model

import (
	"github.com/satori/go.uuid"
	"time"
)

// TODO build the TaskStatus type as an int
// TaskStatus is the current processing status of a task
type TaskStatus int

// TODO define the Task status enum as const using iota (StatusTodo, StatusInProgress, StatusDone)
const (
	StatusTodo TaskStatus = iota
	StatusInProgress
	StatusDone
)

// TODO build the TaskPriority type as an int
// TaskPriority is the priority of a task
type TaskPriority int

// TODO define the Task Priority enum as const using iota (PriorityMinor, PriorityMedium, PriorityHigh)
const (
	PriorityMinor TaskPriority = iota
	PriorityMedium
	PriorityHigh
)

// TODO add the Status and Priority enums, the Creation and Due Dates and the JSON ans BSON annotations

// Task is the structure to define a task to be done
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// TODO Status
	Status TaskStatus `json:"status"`
	// TODO Priority
	Priority TaskPriority `json:"priority"`
	// TODO Creation Date
	CreationDate time.Time `json:"creation_date"`
	// TODO Due Date
	DueDate time.Time `json:"due_date"`
}

// TODO add a NewTask method to create a new UUID for the task when called
// NewTask builds a new task with a new ID of the Task as a string
func NewTask() Task {
	t := Task {
		ID: uuid.NewV4().String(),
		Status: StatusTodo,
		Priority: PriorityMedium,
		CreationDate: time.Now().UTC(),
	}

	return t
}

// TODO add an Equal method for Task comparison, be careful with time.Time comparison
func (t1 Task) Equal(t2 Task) bool {
	return t1 == t2
}
