package main

import "testing"

func TestOpcodeRun(t *testing.T) {
	opcode := NewOpcode([]int{1,0,0,0,99})
	opcode.Run()

	if opcode.Data != [2,0,0,0,99] {
		t.Errorf("for mass %v got %v; want %v", input, output, wantedOutput)
	}
}
