package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOpcodeRun(t *testing.T) {
	input := []int{1, 0, 0, 0, 99}
	wantedOutput := []int{2, 0, 0, 0, 99}
	opcode := NewOpcode(input)
	opcode.Run()
	fmt.Println(opcode)
	output := opcode.Data

	if !reflect.DeepEqual(output, wantedOutput) {
		t.Errorf("for input %v got %v; want %v", input, output, wantedOutput)
	}

	input = []int{2, 3, 0, 3, 99}
	wantedOutput = []int{2, 3, 0, 6, 99}
	opcode = NewOpcode(input)
	opcode.Run()
	fmt.Println(opcode)
	output = opcode.Data

	if !reflect.DeepEqual(output, wantedOutput) {
		t.Errorf("for input %v got %v; want %v", input, output, wantedOutput)
	}

	input = []int{2, 4, 4, 5, 99, 0}
	wantedOutput = []int{2, 4, 4, 5, 99, 9801}
	opcode = NewOpcode(input)
	opcode.Run()
	fmt.Println(opcode)
	output = opcode.Data

	if !reflect.DeepEqual(output, wantedOutput) {
		t.Errorf("for input %v got %v; want %v", input, output, wantedOutput)
	}

	input = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	wantedOutput = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	opcode = NewOpcode(input)
	opcode.Run()
	fmt.Println(opcode)
	output = opcode.Data

	if !reflect.DeepEqual(output, wantedOutput) {
		t.Errorf("for input %v got %v; want %v", input, output, wantedOutput)
	}
}
