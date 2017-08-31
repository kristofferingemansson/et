package et

import (
	"encoding/base64"
	"errors"
	"fmt"
	"testing"
)

var (
	ErrorA = errors.New("ErrorA")
	ErrorB = errors.New("ErrorB")
)

func TestError(t *testing.T) {
	err := func() error {
		_, err := base64.StdEncoding.DecodeString("x")
		if err != nil {
			return New(
				ErrorA,
				err,
			)
		}
		return nil
	}()

	if err != nil {
		fmt.Println(err)
		fmt.Println(Trail(err))
		fmt.Println(Trace(err))

		switch Last(err) {
		case ErrorA:
			fmt.Println("ErrorA occurred")
		case ErrorB:
			fmt.Println("ErrorA occurred")
		}
	}
}
