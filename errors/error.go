package errors

type wrapper struct {
	trail ErrorTrail
	trace StackFrames
}

func (e *wrapper) Error() string {
	if len(e.trail) > 0 {
		return e.trail[0].Error()
	}
	return "Unknown error"
}

// New error
func New(errors ...error) error {
	l := len(errors)
	if l == 0 {
		// No error occurred
		return nil
	}

	last := errors[l-1]
	if te, ok := last.(*wrapper); ok {
		if l == 1 {
			return last
		}
		te.trail = append(errors[0:l-1], te.trail...)
		return last
	}

	return &wrapper{
		trail: errors,
		trace: newTrace(last),
	}
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
