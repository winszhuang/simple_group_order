package utils

import "strconv"

func IsPositiveInteger(str string) bool {
	number, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return number > 0
}
