package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// arrange, act, assert
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{5, 5, 10},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
