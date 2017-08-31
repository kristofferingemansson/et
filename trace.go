package et

import (
	"fmt"
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
