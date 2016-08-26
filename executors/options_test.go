package executors

import (
	"testing"

	"github.com/SierraSoftworks/chieftan-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOptions(t *testing.T) {
	Convey("Options", t, func() {
		action := &models.Action{
			Variables: map[string]string{
				"action":              "Action",
				"actionTask":          "Action",
				"actionConfiguration": "Action",
				"actionVariables":     "Action",
			},
		}

		task := &models.Task{
			Variables: map[string]string{
				"task":              "Task",
				"taskConfiguration": "Task",
				"taskVariables":     "Task",
				"actionTask":        "Task",
			},
		}

		configuration := &models.ActionConfiguration{
			Variables: map[string]string{
				"configuration":          "Configuration",
				"configurationVariables": "Configuration",
				"actionConfiguration":    "Configuration",
				"taskConfiguration":      "Configuration",
				"taskVariables":          "Configuration",
			},
		}

		variables := map[string]string{
			"variables":              "Variables",
			"configurationVariables": "Variables",
			"actionVariables":        "Variables",
			"taskVariables":          "Variables",
		}

		Convey("MergeVariables", func() {
			Convey("With a Configuration", func() {
				options := &Options{
					Action:        action,
					Task:          task,
					Configuration: configuration,
					Variables:     variables,
				}

				v := options.MergeVariables()
				So(v, ShouldNotBeNil)
				So(v["action"], ShouldEqual, "Action")

				So(v["task"], ShouldEqual, "Task")
				So(v["actionTask"], ShouldEqual, "Task")

				So(v["configuration"], ShouldEqual, "Configuration")
				So(v["actionConfiguration"], ShouldEqual, "Configuration")
				So(v["taskConfiguration"], ShouldEqual, "Configuration")

				So(v["variables"], ShouldEqual, "Variables")
				So(v["actionVariables"], ShouldEqual, "Variables")
				So(v["taskVariables"], ShouldEqual, "Variables")
				So(v["configurationVariables"], ShouldEqual, "Variables")
			})
			Convey("Without a Configuration", func() {
				options := &Options{
					Action:    action,
					Task:      task,
					Variables: variables,
				}

				v := options.MergeVariables()
				So(v, ShouldNotBeNil)
				So(v["action"], ShouldEqual, "Action")

				So(v["task"], ShouldEqual, "Task")
				So(v["actionTask"], ShouldEqual, "Task")

				So(v["configuration"], ShouldBeEmpty)
				So(v["actionConfiguration"], ShouldEqual, "Action")
				So(v["taskConfiguration"], ShouldEqual, "Task")

				So(v["variables"], ShouldEqual, "Variables")
				So(v["actionVariables"], ShouldEqual, "Variables")
				So(v["taskVariables"], ShouldEqual, "Variables")
				So(v["configurationVariables"], ShouldEqual, "Variables")
			})
		})
	})
}
