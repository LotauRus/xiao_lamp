package support

import (
	"testing"
)

func TestMapValue(t *testing.T) {
	t.Run("[float32,float32]", func(t *testing.T) {
		testCases := []struct {
			value            uint16
			fromMin, fromMax float32
			toMin, toMax     float32
			expected         float32
		}{
			// Test cases
			{7, 0, 10, 0, 100, 70},
			{2, 2, 10, 0, 100, 0},
			{10, 2, 10, 0, 100, 100},
			{5, 0, 10, -100, 100, 0},
			{15, 0, 10, -100, 100, 100},
			{12, 0, 100, 10, 20, 11.2},
			{18, 0, 100, 10.5, 20.8, 12.354},
		}
		for i, tc := range testCases {
			result := MapValue(tc.value, tc.fromMin, tc.fromMax, tc.toMin, tc.toMax)
			if result != tc.expected {
				t.Errorf("Test case %d failed. Expected %v, but got %v", i, tc.expected, result)
			}
		}
	})

	t.Run("[uint32,float32]", func(t *testing.T) {
		testCases := []struct {
			value            uint16
			fromMin, fromMax uint32
			toMin, toMax     float32
			expected         float32
		}{
			// Test cases
			{7, 0, 10, 0, 100, 70},
			{2, 2, 10, 0, 100, 0},
			{10, 2, 10, 0, 100, 100},
			{12, 0, 100, 10, 20, 11.2},
			{18, 0, 100, 10.5, 20.8, 12.354},
		}
		for i, tc := range testCases {
			result := MapValue(tc.value, tc.fromMin, tc.fromMax, tc.toMin, tc.toMax)
			if result != tc.expected {
				t.Errorf("Test case %d failed. Expected %v, but got %v", i, tc.expected, result)
			}
		}
	})

	t.Run("[uint32,uint32]", func(t *testing.T) {
		testCases := []struct {
			value            uint16
			fromMin, fromMax uint32
			toMin, toMax     uint32
			expected         uint32
		}{
			// Test cases
			{7, 0, 10, 0, 100, 70},
			{2, 2, 10, 0, 100, 0},
			{10, 2, 10, 0, 100, 100},
			{12, 0, 100, 10, 20, 11},
		}
		for i, tc := range testCases {
			result := MapValue(tc.value, tc.fromMin, tc.fromMax, tc.toMin, tc.toMax)
			if result != tc.expected {
				t.Errorf("Test case %d failed. Expected %v, but got %v", i, tc.expected, result)
			}
		}
	})

	t.Run("[float32,uint32]", func(t *testing.T) {
		testCases := []struct {
			value            uint16
			fromMin, fromMax float32
			toMin, toMax     uint32
			expected         uint32
		}{
			// Test cases
			{7, 0, 10, 0, 100, 70},
			{2, 2, 10, 0, 100, 0},
			{10, 2, 10, 0, 100, 100},
			{12, 0, 100, 10, 20, 11},
		}
		for i, tc := range testCases {
			result := MapValue(tc.value, tc.fromMin, tc.fromMax, tc.toMin, tc.toMax)
			if result != tc.expected {
				t.Errorf("Test case %d failed. Expected %v, but got %v", i, tc.expected, result)
			}
		}
	})
}
