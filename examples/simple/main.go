package main

import (
	stderrors "errors"
	"fmt"
	"github.com/kristofferingemansson/go-errors/errors"
)

// Define som package errors
var (
	ErrAuth     = stderrors.New("ErrAuth")
	ErrNotFound = stderrors.New("ErrNotFound")

	ErrReallyBad = stderrors.New("ErrReallyBad")
)

func main() {
	// Create new error with message
	err := errors.New("An error occurred")
	fmt.Println(err) // Print "An error occurred"

	// Lets wrap a package error
	err = errors.New("Invalid username or password", ErrAuth)
	fmt.Println(err)              // Print "Invalid username or password"
	fmt.Println(errors.Last(err)) // Print "ErrAuth"

	// Switch on "last" error
	switch errors.Last(err) {
	case ErrAuth:
		fmt.Println("http 403")
	case ErrNotFound:
		fmt.Println("http 404")
	default:
		fmt.Println("http 500")
	}

	// Lets wrap that error in an other error!
	err = errors.New("", ErrReallyBad, err)

	// Print the "trail" of errors as single line summary
	fmt.Println(errors.Trail(err)) // ErrReallyBad, ErrAuth

	// Print an extended error trail
	fmt.Printf("%+v\n", errors.Trail(err))
	/*
		ErrReallyBad (*errors.errorString: &errors.errorString{s:"ErrReallyBad"})
		ErrAuth (*errors.errorString: &errors.errorString{s:"ErrAuth"})
	*/

	// Print deepest stack trace line
	fmt.Println(errors.Trace(err)) // main.main() @ go-errors/examples/simple/main.go:23

	// Print full stack trace
	fmt.Printf("%+v\n", errors.Trace(err))
	/*
		main.main() @ go-errors/examples/simple/main.go:23
		main.main() @ go-errors/examples/simple/main.go:23
		runtime.main() @ Go/src/runtime/proc.go:185
		runtime.goexit() @ Go/src/runtime/asm_amd64.s:2197
	*/

}
