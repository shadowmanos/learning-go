package ageDistribution

import (
	"testing"
)

func TestProcess_1(t *testing.T) {
	expected := "Still in Mama's arms"
	result := Process(0)
	if result != expected {
		t.Errorf("Process(0) = %q but wanted %q", result, expected)
	}
}

func TestProcess_2(t *testing.T) {
	expected := "College"
	result := Process(19)
	if result != expected {
		t.Errorf("Process(19) = %q but wanted %q", result, expected)
	}
}
