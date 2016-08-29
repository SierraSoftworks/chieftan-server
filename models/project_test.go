package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestProject(t *testing.T) {
	Convey("Project", t, func() {
		Convey("Summary", func() {
			project := Project{
				ID:          "000000000000000000000000",
				Name:        "Test Project",
				Description: "Test",
				URL:         "https://github.com/SierraSoftworks/chieftan-server",
			}

			summary := project.Summary()

			So(summary, ShouldNotBeNil)
			So(summary, ShouldResemble, &ProjectSummary{
				ID:   "000000000000000000000000",
				Name: "Test Project",
				URL:  "https://github.com/SierraSoftworks/chieftan-server",
			})
		})
	})
}
