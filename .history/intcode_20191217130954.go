package main

// Intcode is special computer used by elves
type Intcode struct {
	index int
	Data  []int
}

func (intcode *Intcode) calculate() {
	defer func() {
		r := recover()
		if r != nil {
			// log.Println("recovering")
		}
		intcode.index = intcode.index + 4
	}()
	curentIndex := intcode.index

	param1 := intcode.Data[intcode.Data[curentIndex+1]]
	param2 := intcode.Data[intcode.Data[curentIndex+2]]

	storePosition := intcode.Data[curentIndex+3]

	switch intcode.Data[curentIndex] {
	case 1:
		intcode.Data[storePosition] = param1 + param2
	case 2:
		intcode.Data[storePosition] = param1 * param2
	}
}

// Run starts up Intcode computer and performs calculations
func (intcode *Intcode) Run() {
	for intcode.Data[intcode.index] != 99 {
		intcode.calculate()
	}
}

// NewIntcode return new Intcode computer
func NewIntcode(data []int) *Intcode {
	intcode := &Intcode{index: 0, Data: data}
	return intcode
}
