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
	return len(strings.TrimSpace(v)) > 0
}
