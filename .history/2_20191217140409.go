package main

import "fmt"

func two() {
	var intcode Intcode
	input := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 2, 9, 19, 23, 1, 9, 23, 27, 2, 27, 9, 31, 1, 31, 5, 35, 2, 35, 9, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 47, 6, 51, 2, 51, 10, 55, 1, 9, 55, 59, 2, 6, 59, 63, 1, 63, 6, 67, 1, 67, 10, 71, 1, 71, 10, 75, 2, 9, 75, 79, 1, 5, 79, 83, 2, 9, 83, 87, 1, 87, 9, 91, 2, 91, 13, 95, 1, 95, 9, 99, 1, 99, 6, 103, 2, 103, 6, 107, 1, 107, 5, 111, 1, 13, 111, 115, 2, 115, 6, 119, 1, 119, 5, 123, 1, 2, 123, 127, 1, 6, 127, 0, 99, 2, 14, 0, 0}
	// intcode := NewIntcode(input)
	// input[1] = 12
	// input[2] = 2
	// intcode := NewIntcode(input)
	// intcode.Run()
	// fmt.Println(intcode.Data[0])
	// noun := 0
	// verb := 0
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			newInput := input
			newInput[1] = noun
			newInput[2] = verb
			// fmt.Println(noun, verb, input)
			intcode = *NewIntcode(input)
			intcode.Run()
			// fmt.Println(intcode.Data[0])
			if intcode.Data[0] == 19690720 {
				fmt.Println(noun, verb)
			}
		}
	}
	// intcode := NewIntcode(input)
	// intcode.Run()
	// fmt.Println(intcode.Data[0])
}
