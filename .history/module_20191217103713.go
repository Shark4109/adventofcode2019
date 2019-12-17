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
	requiredFuel := requiredFuelForMass(fuelMass)
	return requiredFuel
}

// NewModule return new instance of Module
func NewModule(mass int) *Module {
	module := &Module{
		Mass: mass,
	}
	return module
}
