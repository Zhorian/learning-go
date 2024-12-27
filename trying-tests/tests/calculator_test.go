package calculator_test

import (
	"testing"
	calculator "trying_tests"
)

func TestAddOf2Positives(t *testing.T) {
	actual := calculator.Add(1, 3)
	expected := 4

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestAddOf1PositiveAnd1Negative(t *testing.T) {
	actual := calculator.Add(7, -3)
	expected := 4

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestAddOf1NegativeAnd1Positive(t *testing.T) {
	actual := calculator.Add(-5, 4)
	expected := -1

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestAddOf2Negatives(t *testing.T) {
	actual := calculator.Add(-2, -6)
	expected := -8

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}
