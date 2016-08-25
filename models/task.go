package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type TaskState int

const TaskStateNotExecuted = 0
const TaskStateExecuting = 1
const TaskStateFailed = 2
const TaskStatePassed = 3

type TaskMetadata struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

type TaskSummary struct {
	ID       bson.ObjectId `json:"id"`
	Metadata TaskMetadata  `json:"metadata"`
}

type Task struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Metadata TaskMetadata  `json:"metadata"`

	Created   time.Time `json:"created"`
	Executed  time.Time `json:"executed"`
	Completed time.Time `json:"completed"`

	Action  *ActionSummary  `json:"action"`
	Project *ProjectSummary `json:"project"`

	Variables map[string]string `json:"vars"`
	State     TaskState         `json:"state"`
	Output    string            `json:"output"`
}

func (t *Task) Summary() *TaskSummary {
	return &TaskSummary{
		ID:       t.ID,
		Metadata: t.Metadata,
	}
}
