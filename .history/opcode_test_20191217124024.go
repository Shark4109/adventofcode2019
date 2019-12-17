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
}
