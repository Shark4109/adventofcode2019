package main

import (
	"fmt"
	"strings"
)

type Wiring struct {
	index       int
	WiringInstructions [][]string
	WiringTable [][]string
}

func (wiring *Wiring) move(code string) {
	if strings.HasPrefix(code, "R") {
		fmt.Println("go right")
	} else if strings.HasPrefix(code, "L") {
	}
}

func NewWiring() *Wiring{
	wiring := 
}
