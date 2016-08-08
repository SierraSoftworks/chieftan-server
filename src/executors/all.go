package executors

var executors map[string]Executor

func Get(operation string) Executor {
	return executors[operation]
}

func Register(operation string, executor Executor) {
	executors[operation] = executor
}
