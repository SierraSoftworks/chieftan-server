package models

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTask(t *testing.T) {
	Convey("Task", t, func() {
		Convey("Summary", func() {
			task := Task{
				ID: "000000000000000000000000",
				Metadata: TaskMetadata{
					Description: "Test task",
				},
				Created: time.Now(),
				Action: ActionSummary{
					ID:          "000000000000000000000000",
					Name:        "Test Action",
					Description: "Test action",
				},
				Project: ProjectSummary{
					ID:   "000000000000000000000000",
					Name: "Test Project",
					URL:  "https://github.com/SierraSoftworks/chieftan-server",
				},
				Variables: map[string]string{
					"x": "1",
				},
				State:  TaskStateNotExecuted,
				Output: "",
			}

			summary := task.Summary()

			So(summary, ShouldNotBeNil)
			So(summary, ShouldResemble, &TaskSummary{
				ID: "000000000000000000000000",
				Metadata: TaskMetadata{
					Description: "Test task",
				},
			})
		})
	})
}
