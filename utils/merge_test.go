package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMerge(t *testing.T) {
	Convey("Merge", t, func() {
		m1 := map[string]string{
			"x1": "1",
			"y":  "1",
		}
		m2 := map[string]string{
			"x2": "2",
			"y":  "2",
		}
		m3 := map[string]string{
			"x2": "3",
			"z":  "3",
		}

		mr := Merge(m1, m2, m3, nil)

		So(mr, ShouldNotBeNil)
		So(mr, ShouldResemble, map[string]string{
			"x1": "1",
			"x2": "3",
			"y":  "2",
			"z":  "3",
		})
	})
}
