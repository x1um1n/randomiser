// tests for functions in randomiser.go
package main

import (
	"testing"
)

// confirm that containsString finds a string in a slice
func TestContainsString(t *testing.T) {
	var tests = []struct {
		slce []string
		in   string
		out  bool
	}{
		{[]string{"one", "two", "three", "four"}, "five", false},
		{[]string{"one", "two", "three", "four"}, "one", true},
	}

	for _, test := range tests {

		if out := containsString(test.slce, test.in); out != test.out {
			t.Errorf("Test Failed: {%s} - {%s} inputted, {%t} expected, recieved: {%t}", test.slce, test.in, test.out, out)
		}
	}
}
