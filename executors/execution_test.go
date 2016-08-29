package executors

import (
	"testing"

	"time"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

type testExecutorProvider struct {
	Executors []Executor
}

func (p *testExecutorProvider) GetExecutors(e *Execution) []Executor {
	if p.Executors == nil {
		return []Executor{}
	}

	return p.Executors
}

type testExecutor struct {
	RunHandler func(ctx *Execution) error
}

func (e *testExecutor) Name() string {
	return "Test"
}

func (e *testExecutor) Run(ctx *Execution) error {
	return e.RunHandler(ctx)
}

func TestExecution(t *testing.T) {
	Convey("Execution", t, func() {

		configuration := &models.ActionConfiguration{
			Name:      "Test",
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

		Convey("Constructor", func() {
			Convey("With no options", func() {
				_, err := NewExecution(nil)
				So(err, ShouldNotBeNil)
			})

			Convey("With no action", func() {
				_, err := NewExecution(&Options{
					Task:          task,
					Configuration: configuration,
					Variables:     variables,
				})
				So(err, ShouldNotBeNil)
			})

			Convey("With no task", func() {
				_, err := NewExecution(&Options{
					Action:        action,
					Configuration: configuration,
					Variables:     variables,
				})
				So(err, ShouldNotBeNil)
			})

			Convey("With no configuration", func() {
				_, err := NewExecution(&Options{
					Action:    action,
					Task:      task,
					Variables: variables,
				})
				So(err, ShouldBeNil)
			})

			Convey("With a configuration", func() {
				_, err := NewExecution(&Options{
					Action:        action,
					Task:          task,
					Configuration: configuration,
					Variables:     variables,
				})
				So(err, ShouldBeNil)
			})
		})

		Convey("With a valid executor", func() {
			ex, err := NewExecution(&Options{
				Action:        action,
				Task:          task,
				Configuration: configuration,
				Variables:     variables,
			})
			So(err, ShouldBeNil)

			hasExecuted := false
			ex.ExecutorProvider = &testExecutorProvider{
				Executors: []Executor{
					&testExecutor{
						RunHandler: func(ctx *Execution) error {
							time.Sleep(100 * time.Millisecond)
							hasExecuted = true
							return nil
						},
					},
				},
			}

			Convey("Should correctly run the executor", func() {
				stateChanged := ex.Start()
				So(stateChanged, ShouldNotBeNil)

				for _ = range stateChanged {

				}

				So(hasExecuted, ShouldBeTrue)
				So(ex.Task.Completed, ShouldNotResemble, time.Time{})
				So(ex.Task.Output, ShouldNotBeEmpty)
				So(ex.Task.State, ShouldEqual, models.TaskStatePassed)
			})
		})
	})
}
