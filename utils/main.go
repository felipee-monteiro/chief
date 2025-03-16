package utils

import (
	"strconv"
	"strings"
)

func IsNumeric(v string) bool {
	_, err := strconv.Atoi(v)

	return err == nil
}

func IsValidString(v string) bool {
	if len(strings.TrimSpace(v)) == 0 {
		return false
	}

	return true
}
