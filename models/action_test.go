package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAction(t *testing.T) {
	Convey("Action", t, func() {
		Convey("Summary", func() {
			action := Action{
				ID:          "000000000000000000000000",
				Name:        "Test Action",
				Description: "Test",
			}

			summary := action.Summary()

			So(summary, ShouldNotBeNil)
			So(summary, ShouldResemble, &ActionSummary{
				ID:          "000000000000000000000000",
				Name:        "Test Action",
				Description: "Test",
			})
		})
	})
}
