package repo

import "testing"

func TestConvert(t *testing.T) {
	numbers := []float64{22.5, 25.5, 25.5}
	expected := 22.5

	var result float64

	result = valute2BTC(numbers[0], numbers[1], numbers[2])
	if result != expected {
		t.Errorf("Expected %f, got %f ", expected, result)
	}

}
