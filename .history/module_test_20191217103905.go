package main

import "testing"

func TestRequiredFuelForMass(t *testing.T) {
	input := 12
	wantedOutput := 2
	output := requiredFuelForMass(input)

	if wantedOutput != output {
		t.Errorf("for mass %v got %v; want %v", input, output, wantedOutput)
	}

	input := 1969
	wantedOutput := 966
	output := requiredFuelForMass(input)

	if wantedOutput != output {
		t.Errorf("for mass %v got %v; want %v", input, output, wantedOutput)
	}
}
