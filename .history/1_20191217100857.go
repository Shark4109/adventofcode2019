package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var modules []*Module

func one() {
	fmt.Print("one")
}

func getModules() {
	file, err := os.OpenFile("modules.txt", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal("could not read file: ", err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
