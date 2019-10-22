package main

import "testing"

func TestSum(t *testing.T) {
	// Given
	tests := map[string]struct {
		x, y     int
		expected int
	}{
		"2 + 3 should be 5": {
			x:        2,
			y:        3,
			expected: 5,
		},
		"-1 + 10 should be 9": {
			x:        -1,
			y:        10,
			expected: 9,
		},
	}

	// When
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := sum(test.x, test.y)

			if result != test.expected {
				t.Errorf("expected %d + %d to be %d but got: %d", test.x, test.y, test.expected, result)
			}
		})
	}

}
