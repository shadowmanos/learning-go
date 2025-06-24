package ageDistribution

import (
	"testing"
)

func TestProcess_1(t *testing.T) {
	expected := "1 2 F 4 B F 7 8 F B"
	result := Process(3, 5, 10)
	if result != expected {
		t.Errorf("Process(3, 5, 10) = %q but wanted %q", result, expected)
	}
}

func TestProcess_2(t *testing.T) {
	expected := "1 F 3 F 5 F B F 9 F 11 F 13 FB 15"
	result := Process(2, 7, 15)
	if result != expected {
		t.Errorf("Process(2, 7, 15) = %q but wanted %q", result, expected)
	}
}
