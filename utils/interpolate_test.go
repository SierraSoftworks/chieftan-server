package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInterpolate(t *testing.T) {
	Convey("Interpolate", t, func() {
		i := NewInterpolator(map[string]string{
			"name": "Bob",
		})

		Convey("Strings", func() {
			r, err := i.Run("Hello {{name}}!")
			So(err, ShouldBeNil)
			So(r, ShouldEqual, "Hello Bob!")
		})

		Convey("Arrays", func() {
			r, err := i.Run([]string{"Hello", "{{name}}"})
			So(err, ShouldBeNil)
			So(r, ShouldResemble, []string{"Hello", "Bob"})
		})

		Convey("Structs", func() {
			r, err := i.Run(struct {
				Name string
			}{Name: "{{name}}"})
			So(err, ShouldBeNil)
			So(r, ShouldResemble, struct {
				Name string
			}{Name: "Bob"})
		})

		Convey("Objects", func() {
			r, err := i.Run(map[string]string{
				"message": "Hello {{name}}!",
			})
			So(err, ShouldBeNil)
			So(r, ShouldResemble, map[string]string{
				"message": "Hello Bob!",
			})
		})

		Convey("DeepObjects", func() {
			r, err := i.Run(map[string]interface{}{
				"message": "Hello {{name}}!",
				"details": map[string]string{
					"name": "{{name}}",
				},
			})
			So(err, ShouldBeNil)
			So(r, ShouldResemble, map[string]interface{}{
				"message": "Hello Bob!",
				"details": map[string]string{
					"name": "Bob",
				},
			})
		})
	})
}
