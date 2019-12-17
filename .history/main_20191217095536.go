package main

import "fmt"

func main() {
	module := NewModule(12)
	fmt.Println(module.RequiredFuel())
	module.Mass = 14
	fmt.Println(module.RequiredFuel())
}
