package utils

import (
	. "gopkg.in/check.v1"
)

type InterpolationSuite struct {
	i *Interpolator
}

var _ = Suite(&InterpolationSuite{
	i: NewInterpolator(map[string]string{
		"name": "Bob",
	}),
})

func (s *InterpolationSuite) StringInterpolation(c *C) {
	c.Assert(s.i, NotNil)

	r, err := s.i.Run("Hello {{name}}!")
	c.Assert(err, IsNil)
	c.Assert(r, Equals, "Hello Bob!")
}

func (s *InterpolationSuite) ObjectInterpolation(c *C) {
	r, err := s.i.Run(map[string]string{
		"message": "Hello {{name}}!",
	})

	c.Assert(err, IsNil)
	c.Assert(r, DeepEquals, map[string]string{
		"message": "Hello Bob!",
	})
}

func (s *InterpolationSuite) DeepObjectInterpolation(c *C) {
	r, err := s.i.Run(map[string]interface{}{
		"message": "Hello {{name}}!",
		"details": map[string]string{
			"name": "{{name}}",
		},
	})

	c.Assert(err, IsNil)
	c.Assert(r, DeepEquals, map[string]interface{}{
		"message": "Hello Bob!",
		"details": map[string]string{
			"name": "Bob",
		},
	})
}
