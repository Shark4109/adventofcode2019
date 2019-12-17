package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func one() {
	fmt.Print("one")
}

func getModules() {
	file, err := os.OpenFile("modules.txt")
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
