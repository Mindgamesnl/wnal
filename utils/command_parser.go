package utils

import (
	"errors"
	"os"
	"strconv"
)

func FindStartingIndex() (error, int) {
	for i := range os.Args {
		arg := os.Args[i]
		if isNumber(arg) {
			return nil, i
		}
	}
	return errors.New("Can't find a starting port"), 0
}

func isNumber(input string) bool {
	if _, err := strconv.Atoi(input); err == nil {
		return true
	}
	return false
}