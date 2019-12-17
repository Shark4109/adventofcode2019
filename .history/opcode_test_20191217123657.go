package main

import "testing"

func TestOpcodeRun(t *testing.T) {
	opcode := NewOpcode([]int{1,0,0,0,99})
	opcode.Run()
	
	if opcode.Data != 
}
