package errors

// wrapper internal error wrapper holding both msg, trail, and trace
type wrapper struct {
	msg   *string
	trail ErrorTrail
	trace StackFrames
}

// Error return string representation of error
func (e *wrapper) Error() string {
	if e.msg != nil {
		return *e.msg
	}

	if len(e.trail) > 0 {
		return e.trail[0].Error()
	}

	return "Unknown error"
}

// New error
// If the last error is of [this package] error type, then the errors will be prepended to its error trail
// Else a new error wrapper will be created with the supplied errors
func New(msg string, errors ...error) error {
	var wrapperError *wrapper

	l := len(errors)
	if l > 0 {
		last := errors[l-1]
		if we, ok := last.(*wrapper); ok {
			we.trail = append(errors[0:l-1], we.trail...)
			wrapperError = we
		} else {
			wrapperError = &wrapper{
				trail: errors,
				trace: newTrace(last),
			}
		}
	} else {
		wrapperError = &wrapper{
			trace: newTrace(nil),
		}
		wrapperError.trail = []error{wrapperError}
	}

	if msg != "" {
		wrapperError.msg = &msg
	}

	return wrapperError
}

// Last returns latest error from trail
func Last(err error) error {
	if err == nil {
		return nil
	}
	if te, ok := err.(*wrapper); ok {
		return te.trail[0]
	}
	return err
}

// First returns first (deepest) error from trail
func First(err error) error {
	if err == nil {
		return nil
	}
	if te, ok := err.(*wrapper); ok {
		return te.trail[len(te.trail)-1]
	}
	return err
}
