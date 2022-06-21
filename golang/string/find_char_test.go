package utils_string

import (
	"testing"
)

const ERROR_DETECTED = "FindChar func with args (char=%s, str=%s), expected to return %d, actual return %d."

func TestFindChar(t *testing.T) {
	char := "a"
	str := "This is a string!"

	expectedResult := 8

	test := FindChar(char, str)

	if test != expectedResult {
		t.Errorf(ERROR_DETECTED, char, str, expectedResult, test)
	}
}

func TestCharNotFound(t *testing.T) {
	char := "à"
	str := "This is a string!"

	expectedResult := -1

	test := FindChar(char, str)

	if test != expectedResult {
		t.Errorf(ERROR_DETECTED, char, str, expectedResult, test)
	}
}
