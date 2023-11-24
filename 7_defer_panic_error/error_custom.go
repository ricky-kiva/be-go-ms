package main

import (
	"errors"
	"strings"
)

func validateFilledString(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("String cannot be blank")
	}
	return true, nil
}
