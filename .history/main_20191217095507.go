package main

import "fmt"

func main() {
	module := NewModule(12)
	fmt.Println(module.RequiredFuel())
}
