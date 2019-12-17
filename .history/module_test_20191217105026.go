package main

import "testing"

func TestRequiredFuelForMass(t *testing.T) {
	input := 12
	wantedOutput := 2
	output := requiredFuelForMass(input)

	if wantedOutput != output {
		t.Errorf("for mass %v got %v; want %v", input, output, wantedOutput)
	}

	input = 1969
	wantedOutput = 654
	output = requiredFuelForMass(input)

	if wantedOutput != output {
		t.Errorf("for mass %v got %v; want %v", input, output, wantedOutput)
	}
}

func TestRequiredFuelForFuel(t *testing.T) {
	input := 2
	wantedOutput := 0
	output := requiredFuelForFuel(input)

	if wantedOutput != output {
		t.Errorf("for mass %v got %v; want %v", input, output, wantedOutput)
	}

	input = 654
	wantedOutput = 312
	output = requiredFuelForFuel(input)

	if wantedOutput != output {
		t.Errorf("for mass %v got %v; want %v", input, output, wantedOutput)
	}
}

func TestCalculateRequiredFuel(t *testing.T) {
	input := 12
	wantedOutput := 2
	module := NewModule(input)

	if module.RequiredFuel != wantedOutput {
		t.Errorf("for mass %v got %v; want %v", input, module.RequiredFuel, wantedOutput)
	}

	input = 1969
	wantedOutput = 966
	module = NewModule(input)

	if module.RequiredFuel != wantedOutput {
		t.Errorf("for mass %v got %v; want %v", input, module.RequiredFuel, wantedOutput)
	}
}
