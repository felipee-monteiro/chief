package utils

import (
	"os"
	"strconv"
	"strings"
)

func IsNumeric(v string) bool {
	_, err := strconv.Atoi(v)
	return err == nil
}

func IsValidString(v string) bool {
	return len(strings.TrimSpace(v)) == 0
}

func IsDirValid(stat os.FileInfo) bool {
	return stat.IsDir()
}
