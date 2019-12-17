package main

import (
	"fmt"
	"strings"
)

type Wiring struct {
	index       int
	WiringTable [][]int
}

func (wiring *Wiring) move(code string) {
	if strings.HasPrefix(code, "R") {
		fmt.Println("go right")
	} else if strings.HasPrefix(code, "L") {
	}
}
