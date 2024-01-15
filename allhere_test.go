package main

import "testing"

func TestCodeRange(t *testing.T) {
	got := isValidStatusCodeRange(250)

	if got == false {
		t.Errorf("Invalid code range")
	}
}

func TestStringToInt(t *testing.T) {
	_, err := stringToInt("250")

	if err != nil {
		t.Errorf("Invalid code range")
	}
}
