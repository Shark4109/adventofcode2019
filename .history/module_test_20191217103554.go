package main

import "testing"

func TestRequiredFuelForMass(t *testing.T) {
	input := 12
	wantedOutput := 2
	output := requiredFuelForMass(input)

	if wantedOutput != output {
		t.Errorf("got %v; want %v", output, input)
	}
}
