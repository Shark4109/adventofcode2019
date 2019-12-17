package main

// Module to launch
type Module struct {
	Mass         int
	RequiredFuel int
}

// CalculateRequiredFuel return how much fuled does module need to lift off
func (module *Module) calculateRequiredFuel() {
	requiredFuelForModuleMass := requiredFuelForMass(module.Mass)
	module.RequiredFuel = requiredFuelForModuleMass + requiredFuelForFuel(requiredFuelForModuleMass)
}

func requiredFuelForMass(mass int) int {
	return int(mass/3) - 2
}

func requiredFuelForFuel(fuelMass int) int {
	totalRequiredFuel := 0
	requiredFuel := requiredFuelForMass(fuelMass)
	for requiredFuel > 0 {
		totalRequiredFuel += requiredFuel
		fuelMass = requiredFuel
		requiredFuel := requiredFuelForMass(fuelMass)
	}
	return totalRequiredFuel
}

// NewModule return new instance of Module
func NewModule(mass int) *Module {
	module := &Module{
		Mass: mass,
	}
	return module
}
