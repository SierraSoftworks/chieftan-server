package utils

import (
	"log"

	"github.com/aymerick/raymond"
)

type Interpolator struct {
	Context map[string]string
}

func NewInterpolator(context map[string]string) *Interpolator {
	return &Interpolator{
		Context: context,
	}
}

func (e *Interpolator) Run(item interface{}) (interface{}, error) {
	switch item.(type) {
	case string:
		return e.interpolateString(item.(string))

	case map[string]interface{}:
		return e.interpolateObject(item.(map[string]interface{}))

	case map[string]string:
		return e.interpolateFlatMap(item.(map[string]string))

	default:
		log.Printf("Failed to interpolate %#v\n", item)
		return item, nil
	}
}

func (e *Interpolator) interpolateString(item string) (string, error) {
	return raymond.Render(item, e.Context)
}

func (e *Interpolator) interpolateFlatMap(item map[string]string) (map[string]string, error) {
	result := make(map[string]string)

	for k, v := range item {
		val, err := e.interpolateString(v)
		if err != nil {
			return nil, err
		}
		result[k] = val
	}

	return result, nil
}

func (e *Interpolator) interpolateObject(item map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for k, v := range item {
		val, err := e.Run(v)
		if err != nil {
			return nil, err
		}
		result[k] = val
	}

	return result, nil
}
