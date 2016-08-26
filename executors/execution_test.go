package executors

import (
	"testing"

	"time"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

type testExecution struct {
	*Execution

	Executors []Executor
}

func (e *testExecution) GetExecutors() []Executor {
	return e.Executors
}

func TestExecution(t *testing.T) {
	Convey("Execution", t, func() {

		configuration := &models.ActionConfiguration{
			Variables: map[string]string{},
		}

		action := &models.Action{
			Variables: map[string]string{},
			Configurations: []models.ActionConfiguration{
				*configuration,
			},
		}

		task := &models.Task{
			Variables: map[string]string{},
			Action:    action.Summary(),
		}

		variables := map[string]string{}

		exBase, err := NewExecution(&Options{
			Action:        action,
			Task:          task,
			Configuration: configuration,
			Variables:     variables,
		})

		So(err, ShouldBeNil)
		ex := &testExecution{
			Execution: exBase,
			Executors: []Executor{
				&testExecutor{
					RunHandler: func(ctx *Execution) error {
						time.Sleep(100 * time.Millisecond)
						return nil
					},
				},
			},
		}

		So(ex.GetExecutors(), ShouldResemble, ex.Executors)

		stateChanged := ex.Start()
		So(stateChanged, ShouldNotBeNil)

		for range stateChanged {
		}

		So(ex.Task.State, ShouldEqual, models.TaskStatePassed)
	})
}
