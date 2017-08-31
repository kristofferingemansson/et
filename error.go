package et

type er struct {
	trail ErrorTrail
	trace StackTrace
}

func (e *er) Error() string {
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
	if te, ok := last.(*er); ok {
		if l == 1 {
			return last
		}
		te.trail = append(errors[0:l-1], te.trail...)
		return last
	}

	return &er{
		trail: errors,
		trace: newTrace(last),
	}
}

// Last returns latest error from trail
func Last(err error) error {
	if err == nil {
		return nil
	}
	if te, ok := err.(*er); ok {
		return te.trail[0]
	}
	return err
}
