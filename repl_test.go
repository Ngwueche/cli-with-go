package main

import "testing"

// TestCleanInput demonstrates a table-driven test pattern.
func TestCleanInput(t *testing.T) {
	// Each test case is an anonymous struct with input + expected output.
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	// Range over the test cases and run the same logic for each.
	for _, cs := range cases {
		actual := cleanInput(cs.input)

		// Validate slice lengths before comparing elements.
		if len(actual) != len(cs.expected) {
			t.Errorf("Lengths are not equal: %v vs %v",
				len(actual),
				len(cs.expected),
			)
			continue
		}

		// Compare each element; range over indices for parallel access.
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v does not equal the value of %v",
					expectedWord,
					actualWord,
				)
			}
		}
	}
}
