package et

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
)

// StackTrace ..
type StackTrace []uintptr

// Trace returns stack trace for error
func Trace(err error) StackTrace {
	if te, ok := err.(*er); ok {
		return te.trace
	}

	return nil
}

// Format format StackTrace as string
func (t StackTrace) Format(f fmt.State, c rune) {
	flag := f.Flag('+')
	frames := runtime.CallersFrames(t)
	sep := ""
	for {
		frame, more := frames.Next()
		fmt.Fprintf(f, "%v%v() @ %v:%v", sep, frame.Function, frame.File, frame.Line)
		if !more || !flag {
			break
		}
		sep = "\n"
	}
}

func newTrace(err error) StackTrace {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	// Compatibility with github.com/pkg/errors
	if st, ok := err.(stackTracer); ok {
		tr := st.StackTrace()
		ret := make(StackTrace, len(tr))
		for i, f := range tr {
			ret[i] = uintptr(f)
		}
		return ret
	}

	traceBuf := make(StackTrace, 32)
	n := runtime.Callers(3, traceBuf)
	return traceBuf[0:n]
}
