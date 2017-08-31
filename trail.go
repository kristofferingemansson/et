package et

// ErrorTrail error trail
type ErrorTrail []error

// Trail returns trail of errors
func Trail(err error) ErrorTrail {
	if te, ok := err.(*er); ok {
		return te.trail
	}
	return ErrorTrail{err}
}
