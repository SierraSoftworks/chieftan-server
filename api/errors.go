package api

import (
	"github.com/SierraSoftworks/girder/errors"
	raven "github.com/getsentry/raven-go"
)

type ErrorFormatter struct {
	DefaultFormatter errors.ErrorFormatter
}

func (f *ErrorFormatter) Format(err error) *errors.Error {
	e := f.DefaultFormatter.Format(err)

	if e.Code == 500 {
		trace := e.Stacktrace.(*raven.Stacktrace)

		packet := raven.NewPacket(err.Error())
		packet.Level = raven.ERROR
		packet.Interfaces = append(
			packet.Interfaces,
			raven.NewException(err, trace),
		)

		raven.Capture(packet, nil)
	}

	return e
}

func init() {
	errors.Formatter = &ErrorFormatter{
		DefaultFormatter: errors.Formatter,
	}
}

type StacktraceProvider struct {
}

func (p *StacktraceProvider) Get() interface{} {
	return raven.NewStacktrace(2, 3, []string{"github.com/SierraSoftworks/chieftan-server"})
}
