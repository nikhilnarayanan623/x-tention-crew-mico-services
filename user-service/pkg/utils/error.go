package utils

import (
	"errors"
	"fmt"
)

// To prepend the message to the error
func PrependMessageToError(err error, message string) error {
	if err == nil {
		return errors.New(message)
	}
	return fmt.Errorf("%s \n%w", message, err)
}
