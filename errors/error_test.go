package errors

import (
	"encoding/base64"
	"errors"
	"fmt"
	"testing"
)

var (
	ErrorA = errors.New("Domain error 1")
	ErrorB = errors.New("Domain error 2")
)

func TestError(t *testing.T) {
	err := func() error {
		_, err := base64.StdEncoding.DecodeString("x")
		if err != nil {
			return New(
				"",
				ErrorA,
				errors.New("Description of why everyting went wrong"),
				err,
			)
		}
		return nil
	}()

	if err != nil {
		fmt.Println(err)
		fmt.Println("-------")
		fmt.Println(Trail(err))
		fmt.Println("-------")
		fmt.Printf("%+v\n", Trail(err))
		fmt.Println("-------")
		fmt.Println(Trace(err))
		fmt.Println("-------")
		fmt.Printf("%+v\n", Trace(err))
		fmt.Println("-------")

		switch Last(err) {
		case ErrorA:
			fmt.Println("ErrorA occurred")
		case ErrorB:
			fmt.Println("ErrorA occurred")
		}
	}
}
