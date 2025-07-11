package main

import "testing"

func TestAverage(t *testing.T) {
	grades := map[string]float64{
		"Math":    90,
		"Physics": 80,
		"English": 100,
	}

	expected := 90.0
	result := average(grades)

	if result != expected {
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}
}
