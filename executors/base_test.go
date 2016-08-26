package executors

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testExecutor struct {
	RunHandler func(ctx *Execution) error
}

func (e *testExecutor) Name() string {
	return "Test"
}

func (e *testExecutor) Run(ctx *Execution) error {
	return e.RunHandler(ctx)
}

func TestExecutorBase(t *testing.T) {
	Convey("ExecutorBase", t, func() {
		Convey("Default Implementation", func() {
			e := &ExecutorBase{}

			So(e.Name(), ShouldEqual, "not implemented")
			So(e.Run(nil), ShouldNotBeNil)
			So(e.Run(nil).Error(), ShouldEqual, "not implemented")
		})

		Convey("Custom Implementation", func() {
			var executor Executor
			executor = &testExecutor{
				RunHandler: func(ctx *Execution) error {
					return nil
				},
			}

			So(executor.Name(), ShouldEqual, "Test")
			So(executor.Run(nil), ShouldBeNil)
		})
	})
}
