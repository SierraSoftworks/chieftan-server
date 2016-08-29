package utils

import (
	"reflect"

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
	v := reflect.ValueOf(item)
	newV, err := e.interpolateReflectedValue(v)
	if err != nil {
		return nil, err
	}

	return newV.Interface(), nil
}

func (e *Interpolator) interpolateReflectedValue(v reflect.Value) (*reflect.Value, error) {
	switch v.Kind() {
	case reflect.Interface:
		val, err := e.interpolateReflectedValue(v.Elem())
		if err != nil {
			return nil, err
		}

		out := reflect.ValueOf(val.Interface())
		return &out, nil
	case reflect.Struct:
		out := reflect.New(v.Type()).Elem()
		n := v.NumField()
		for i := 0; i < n; i++ {
			newVal, err := e.interpolateReflectedValue(v.Field(i))
			if err != nil {
				return nil, err
			}

			if newVal != nil {
				out.Field(i).Set(*newVal)
			}
		}

		return &out, nil

	case reflect.Map:
		out := reflect.MakeMap(v.Type())
		for _, key := range v.MapKeys() {
			val, err := e.interpolateReflectedValue(v.MapIndex(key))
			if err != nil {
				return nil, err
			}

			out.SetMapIndex(key, *val)
		}

		return &out, nil

	case reflect.Slice:
		if v.IsNil() {
			return &v, nil
		}

		out := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
		n := v.Len()
		for i := 0; i < n; i++ {
			val := v.Index(i)
			newVal, err := e.interpolateReflectedValue(val)
			if err != nil {
				return nil, err
			}

			if newVal != nil {
				out.Index(i).Set(*newVal)
			}
		}

		return &out, nil

	case reflect.Array:
		out := reflect.New(v.Type())
		n := v.Len()
		for i := 0; i < n; i++ {
			val := v.Index(i)
			newVal, err := e.interpolateReflectedValue(val)
			if err != nil {
				return nil, err
			}

			if newVal != nil {
				out.Index(i).Set(*newVal)
			}
		}

		return &out, nil

	case reflect.String:
		iStr, err := e.interpolateString(v.String())
		if err != nil {
			return nil, err
		}

		out := reflect.ValueOf(iStr)
		return &out, nil
	default:
		return &v, nil
	}
}

func (e *Interpolator) interpolateString(item string) (string, error) {
	return raymond.Render(item, e.Context)
}
