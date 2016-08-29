package executors

import (
	"fmt"
	"time"

	"github.com/SierraSoftworks/chieftan-server/models"
	"github.com/SierraSoftworks/girder/errors"
)

type Execution struct {
	Task          *models.Task
	Action        *models.Action
	Configuration *models.ActionConfiguration
	Options       *Options
	Variables     map[string]string

	ExecutorProvider ExecutorProvider
	StateChanged     chan *models.Task
}

type Executor interface {
	Name() string
	Run(ctx *Execution) error
}

func NewExecution(options *Options) (*Execution, error) {
	if options == nil {
		return nil, errors.BadRequest()
	}

	if options.Action == nil {
		return nil, errors.BadRequest()
	}

	if options.Task == nil {
		return nil, errors.BadRequest()
	}

	return &Execution{
		Task:             options.Task,
		Action:           options.Action,
		Configuration:    options.Configuration,
		Options:          options,
		Variables:        options.MergeVariables(),
		ExecutorProvider: &DefaultExecutorProvider{},
	}, nil
}

func (e *Execution) Start() <-chan *models.Task {
	e.StateChanged = make(chan *models.Task)
	go func() {
		e.Task.Executed = time.Now()
		e.Task.State = models.TaskStateExecuting

		if e.Configuration != nil {
			e.WriteInfo("msg=\"Running task\" configuration=\"%s\"", e.Options.Configuration.Name)
		} else {
			e.WriteInfo("msg=\"Running task\" configuration=\"default\"")
		}
		e.PublishChanges()

		executedSuccessfully := true
		for _, executor := range e.ExecutorProvider.GetExecutors(e) {
			startTime := time.Now()
			e.WriteInfo("msg=\"Starting\" executor=\"%s\"", executor.Name())
			err := executor.Run(e)
			if err != nil {
				e.WriteError("msg=\"Failed\" executor=\"%s\" duration=\"%dms\" error=\"%s\"", executor.Name(), time.Now().Sub(startTime).Nanoseconds()/1e6, err.Error())
				executedSuccessfully = false
				break
			} else {
				e.WriteInfo("msg=\"Completed\" executor=\"%s\" duration=\"%dms\"", executor.Name(), time.Now().Sub(startTime).Nanoseconds()/1e6)
			}

			e.PublishChanges()
		}

		if executedSuccessfully {
			e.WriteInfo("msg=\"Completed\" duration=\"%dms\"", time.Now().Sub(e.Task.Executed).Nanoseconds()/1e6)
			e.Task.State = models.TaskStatePassed
		} else {
			e.WriteError("msg=\"Failed\" duration=\"%dms\"", time.Now().Sub(e.Task.Executed).Nanoseconds()/1e6)
			e.Task.State = models.TaskStateFailed
		}

		e.Task.Completed = time.Now()
		e.PublishChanges()

		close(e.StateChanged)
	}()

	return e.StateChanged
}

func (e *Execution) Write(message string, args ...interface{}) {
	e.Task.Output = fmt.Sprintf("%s%s", e.Task.Output, fmt.Sprintf(message, args...))
}

func (e *Execution) WriteLn(message string, args ...interface{}) {
	e.Task.Output = fmt.Sprintf("%s%s\n", e.Task.Output, fmt.Sprintf(message, args...))
}

func (e *Execution) writeMessage(level, message string, args ...interface{}) {
	e.Task.Output = fmt.Sprintf("%s::[%s] %s::\n", e.Task.Output, level, fmt.Sprintf(message, args...))
}

func (e *Execution) WriteInfo(message string, args ...interface{}) {
	e.writeMessage("info", message, args...)
}

func (e *Execution) WriteError(message string, args ...interface{}) {
	e.writeMessage("error", message, args...)
}

func (e *Execution) PublishChanges() {
	e.StateChanged <- e.Task
}
