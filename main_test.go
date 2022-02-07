package main

import (
	"os"
	"testing"
)

func Test_filesExist(t *testing.T) {
	tests := []struct {
		input string
	}{
		{input: "./main.go"},
		{input: "./go.sum"},
		{input: "./go.mod"},
	}
	for _, testCase := range tests {
		got, err := os.Stat(testCase.input)

		if err != nil {
			t.Errorf("expected %s, got %s", testCase.input, got)
		}
	}
}
