# et
Error+trace for Go

The goal of this package is to provide, and separate, versatile error handling with stack traces in Go.

The package consists of an error wrapper with a simple error stack, and a call stack trace of the first error.
Nested wrapping will only push new errors on top of the error stack. This will allow for easy retrieval and traversal of the error stack.
Adding custom per-package "domain" errors on top of actual errors will allow easy error handling in consuming packages.

## Usage example

```
import (
	"https://github.com/kristofferingemansson/go-error-trace/et"
	"errors"
	"fmt"
)

func main() {
	err := DoStuff()
	if err != nil {
		fmt.Println(err) // Will print either "X occurred" or "Y occurred"
		
		switch et.Last(err) {
		case DoStuffError1:
			// ...
		case DoStuffError1:
			// ...
		default:
			// ...
		}
	}
}

var (
	DoStuffError1 = errors.New("X occurred")
	DoStuffError2 = errors.New("Y occurred")
)

func DoStuff() error {
	if err := OtherFunctionProducingErrorOfUnknownType(); err != nil {
		return et.New(DoStuffError1, err)
	}
	if err := Anotherfunction(); err != nil {
		return et.New(DoStuffError2, err)
	}
	return nil
}
```

## Nesting

A major deal with this package is to wrap(wrap(wrap(your errors))) so that you each time add per-package "domain" errors that can easily be switched on.

```

func main() {
	if err := y.Y(); err != nil {
		switch et.Last(err) {
		case y.ErrYa:
			// Ignore this error
		case y.ErrYb:
			// Print full stack trace
			fmt.Printf("%+v", et.Trace(err))
		}
	}
}

// PackageY
var (
	ErrYa = errors.New("ErrYa")
	ErrYb = errors.New("ErrYb")
)

func Y() error {
	if err := x.X(); err != nil {
		switch et.Last(err) {
		case x.ErrXa, x.ErrXb:
			return et.New(ErrYa, err)
		default:
			return et.New(ErrYb, errors.New("Unhandled error"), err)
		}
	}
	return nil
}


// PackageX
var (
	ErrXa = errors.New("ErrXa")
	ErrXb = errors.New("ErrXb")
)
func Y() error {
	if err := ???; err != nil {
		return et.New(ErrXa, err)
	}
	if err := ???; err != nil {
		return et.New(ErrXb, errors.New("Oh noes!"), err)
	}
	retur nil
}
```

## Error stack
### Printing
```
err := et.New(errors.New("Third"), errors.New("Second"), errors.New("First"))
fmt.Println(err) // "Third"
fmt.Println(et.Trail(err)) // Third, Second, First
fmt.Printf("%v", et.Trail(err)) // (same as above)
fmt.Printf("%+v", et.Trail(err)) // Newline-separated full description of every error object in the stack
```
### Iterating
```
for _, e := range et.Trail(err) {
	// Do stuff with each error in the stack
}
```

## Stack trace
### Printing
```
fmt.Println(et.Trace(err)) // Print deepest stack point on single line
fmt.Printf("%v", et.Trace)) // (same as above)
fmt.Printf("%+v", et.Trace)) // Print newline-separated full stack trace
```
### Iterating
```
frames := runtime.CallersFrames(et.Trace(err))
// Do custom handling of stack frames
```

# Bonus!
Compatibility with `github.com/pkg/errors:`

If you have a package which uses above errors package, the stack trace of an error will be transferred when using `et.New(/*..., */ err)`


