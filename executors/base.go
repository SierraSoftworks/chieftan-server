package executors

import (
	"fmt"
	"time"

	"github.com/SierraSoftworks/chieftan-server/models"
)

type Executor interface {
	Start(ctx *Context) error
}

type ExecutorBase struct {
}

type Context struct {
	Action    *models.Action
	Task      *models.Task
	Variables map[string]string
}

func (e *ExecutorBase) run(ctx *Context) error {
	return fmt.Errorf("not implemented")
}

func (e *ExecutorBase) Start(ctx *Context) error {
	ctx.Task.Executed = time.Now()
	ctx.Task.Output = "::[info] Running task::"
	ctx.Task.State = models.TaskStateExecuting

	if err := models.DB().Tasks().UpdateId(ctx.Task.ID, ctx.Task); err != nil {
		return err
	}

	if err := e.run(ctx); err != nil {
		ctx.Task.Output = fmt.Sprintf("%s\n::[error] Task failed in %dms::", ctx.Task.Output, time.Now().Sub(ctx.Task.Executed).Nanoseconds()/1e6)
		ctx.Task.Output = fmt.Sprintf("%s\n%s", ctx.Task.Output, err.Error())
		ctx.Task.State = models.TaskStateFailed
	} else {
		ctx.Task.Output = fmt.Sprintf("%s\n::[info] Task complete in %dms::", ctx.Task.Output, time.Now().Sub(ctx.Task.Executed).Nanoseconds()/1e6)
		ctx.Task.State = models.TaskStatePassed
	}

	ctx.Task.Completed = time.Now()

	if err := models.DB().Tasks().UpdateId(ctx.Task.ID, ctx.Task); err != nil {
		return err
	}

	return nil
}
