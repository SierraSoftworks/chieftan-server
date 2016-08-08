package tasks

import (
	"fmt"
)

type TaskError struct {
	Code    int    `json:"code"`
	Name    string `json:"error"`
	Message string `json:"message"`
}

func (t *TaskError) Error() string {
	return fmt.Sprintf("%d %s: %s", t.Code, t.Name, t.Message)
}

func NewError(code int, name string, message string) *TaskError {
	return &TaskError{
		Code:    code,
		Name:    name,
		Message: message,
	}
}
