package executors

type ExecutorProvider interface {
	GetExecutors(execution *Execution) []Executor
}

type DefaultExecutorProvider struct {
}

func (p *DefaultExecutorProvider) GetExecutors(e *Execution) []Executor {
	executors := []Executor{}
	if e.Action.HTTP != nil {
		executors = append(executors, &HTTP{})
	}

	return executors
}
