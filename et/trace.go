package et

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
)

// StackFrames ..
type StackFrames []uintptr

// Frames returns stack frames for usage in other errors / custom stack tracing
func Frames(err error) StackFrames {
	if te, ok := err.(*errorWrapper); ok {
		return te.trace
	}
	return nil
}

// StackTrace ..
type StackTrace []string

// Trace returns stack trace for error
func Trace(err error) StackTrace {
	if te, ok := err.(*errorWrapper); ok {
		return stackFramesToStackTrace(te.trace)
	}
	return nil
}

func stackFramesToStackTrace(f StackFrames) StackTrace {
	tr := make(StackTrace, 0)
	frames := runtime.CallersFrames(f)
	for {
		frame, more := frames.Next()
		tr = append(
			tr,
			fmt.Sprintf("%v() @ %v:%v", frame.Function, frame.File, frame.Line),
		)
		if !more {
			break
		}
	}
	return tr
}

// Format format StackTrace as string
func (t StackTrace) Format(f fmt.State, c rune) {
	flag := f.Flag('+')
	for i, tr := range t {
		if i > 0 {
			fmt.Fprint(f, "\n")
		}
		fmt.Fprint(f, tr)
		if !flag {
			break
		}
	}
}

func newTrace(err error) StackFrames {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	// Compatibility with github.com/pkg/errors
	if st, ok := err.(stackTracer); ok {
		tr := st.StackTrace()
		ret := make(StackFrames, len(tr))
		for i, f := range tr {
			ret[i] = uintptr(f)
		}
		return ret
	}

	traceBuf := make(StackFrames, 32)
	n := runtime.Callers(3, traceBuf)
	return traceBuf[0:n]
}
