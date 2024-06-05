package main

import (
	"testing"
)

type testCronField struct {
	data          string
	allowedValues []string
	expected      string
	expectError   bool
}

func TestCronField(t *testing.T) {
	tests := []testCronField{
		{"*", minuteValues, "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59", false},
		{"*/15", minuteValues, "0 15 30 45", false},
		{"1,15", minuteValues, "1 15", false},
		{"5-10", minuteValues, "5 6 7 8 9 10", false},
		{"?", weekValues, "any", false},
		{"5#3", weekValues, "the 3th 5", false},
		{"60", minuteValues, "60", true},           // out of range
		{"invalid", minuteValues, "invalid", true}, // invalid input
		{"1-5/2", minuteValues, "1 3 5", false},
		{"*/2", hourValues, "0 2 4 6 8 10 12 14 16 18 20 22", false},
		{"10-20/3", hourValues, "10 13 16 19", false},
	}
	for _, test := range tests {
		result, err := cronField(test.data, test.allowedValues)
		if (err != nil) != test.expectError {
			t.Errorf("For field %s, expected %s, got %s", test.data, test.expected, result)
		}

		if result != test.expected {
			t.Errorf("For field %s, expected %s, got %s", test.data, test.expected, result)

		}
	}
}

func TestValidateNumber(t *testing.T) {
	tests := []testCronField{
		{"0", minuteValues, "0", false},
		{"59", minuteValues, "59", false},
		{"60", minuteValues, " ", true},
		{"abc", minuteValues, " ", true},
	}

	for _, test := range tests {
		result, err := validateNumber(test.data, test.allowedValues)
		if (err != nil) != test.expectError {
			t.Errorf("For field %s, expected %s, got %s", test.data, test.expected, result)
		}
		if result != test.expected {
			t.Errorf("For field %s, expected %s, got %s", test.data, test.expected, result)
		}
	}
}
