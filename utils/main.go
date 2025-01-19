package utils

import (
    "strconv"
)

func IsNumeric(v string) bool {
    _, err := strconv.Atoi(v)

    return err == nil
}
