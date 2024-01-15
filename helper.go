package main

import (
	"strconv"
)

func stringToInt(value string) (int, error) {
	/**
	This function will convert string to int.
	@var value string.
	@return (int, error).
	*/
	i, err := strconv.Atoi(value)

	if (err) != nil {
		return 0, err
	}

	return i, nil

}

func isValidStatusCodeRange(code int) bool {
	/**
	This function validate the range of the code.
	@var code int.
	@return bool.
	*/
	if code >= 100 && code <= 599 {
		return true
	}

	return false
}
