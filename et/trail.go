package et

import (
	"fmt"
)

// ErrorTrail error trail
type ErrorTrail []error

// Trail returns trail of errors
func Trail(err error) ErrorTrail {
	if te, ok := err.(*er); ok {
		return te.trail
	}
	return ErrorTrail{err}
}

// Format format ErrorTrail as string
func (t ErrorTrail) Format(f fmt.State, c rune) {
	separator := ", "
	format := "%v"

	if f.Flag('+') {
		format = "%[1]v (%[1]T: %#[1]v)"
		separator = "\n"
	}

	for i, err := range t {
		if i > 0 {
			fmt.Fprint(f, separator)
		}
		fmt.Fprintf(f, format, err)
	}
}
