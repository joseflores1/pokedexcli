package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {

	type testCase struct {
		input string
		expected []string
	}

	cases := []testCase{
		{
			input: "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input: "     I LIKe big fruits and I CANNOT lIE   ",
			expected: []string{"i", "like", "big", "fruits", "and", "i", "cannot", "lie"},
		},
		{
			input: "",
			expected: []string{""},
		},
		{
			input: "                         ",
			expected: []string{""},
		},
	}

	passCount := 0
	failCount := 0
	passed := true

	for j, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			failCount++
			t.Errorf(`----------------------------------
Error: Different Lenghts
Input: %v
Expecting: %v
Expecting length: %v
Actual: %v
Actual length: %v
`, c.input, c.expected, len(c.expected), actual, len(actual))
			continue
		} 
		for i := range actual {
			if actual[i] != c.expected[i] {
				failCount++
				passed = false
		t.Errorf(`----------------------------------
Error: Different elements at index %d
Input: %v
Expecting: %v
Actual: %v
`, i, c.input, c.expected, actual)
				break
			}
		}
	
		if passed {
			passCount ++
			fmt.Printf(`----------------------------------
Test %d passed!
Test
Input: %v
Expecting: %v
Actual: %v
`, j, c.input, c.expected, actual)
		}
	}

fmt.Println("----------------------------------")
fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}