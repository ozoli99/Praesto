package utils

import (
	"fmt"
)

func WrapError(err error, context string) error {
	return fmt.Errorf("%s: %w", context, err)
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}