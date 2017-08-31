package et

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
)

type trace []uintptr

// StackTrace pretty printable stack trace
type StackTrace []string

// Trace returns stack trace for error
func Trace(err error) StackTrace {
	if te, ok := err.(*er); ok {
		return traceToStackTrace(te.trace)
	}

	return nil
}

func traceToStackTrace(t trace) StackTrace {
	ret := make(StackTrace, 0)
	frames := runtime.CallersFrames(t)
	for {
		frame, more := frames.Next()
		ret = append(
			ret,
			fmt.Sprintf("%v() @ %v:%v", frame.Function, frame.File, frame.Line),
		)
		if !more {
			break
		}
	}
	return ret
}

func newTrace(err error) trace {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	// Compatibility with github.com/pkg/errors
	if st, ok := err.(stackTracer); ok {
		tr := st.StackTrace()
		ret := make(trace, len(tr))
		for i, f := range tr {
			ret[i] = uintptr(f)
		}
		return ret
	}

	traceBuf := make(trace, 32)
	n := runtime.Callers(3, traceBuf)
	return traceBuf[0:n]
}
