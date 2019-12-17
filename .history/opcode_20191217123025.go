package main

type Opcode struct {
	index int
	Data  []int
}

func (opcode *Opcode) calculate() {
	curentIndex := opcode.index

	param1 := opcode.Data[curentIndex+1]
	param2 := opcode.Data[curentIndex+2]

	storePosition := opcode.Data[curentIndex+3]

	switch opcode.Data[curentIndex] {
	case 1:
		opcode.Data[storePosition] = param1 + param2
	}
}

func (opcode *Opcode) Run() {

}
