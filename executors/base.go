package executors

import "fmt"

type Executor interface {
	Run(ctx *Execution) error
	Name() string
}

type ExecutorBase struct {
}

func (e *ExecutorBase) Name() string {
	return "not implemented"
}

func (e *ExecutorBase) Run(ctx *Execution) error {
	return fmt.Errorf("not implemented")
}
