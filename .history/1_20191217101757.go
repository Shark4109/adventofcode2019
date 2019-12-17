package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var modules []*Module

func one() {
	getModules()
}

func getModules() {
	file, err := os.OpenFile("modules.txt", os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal("could not read file: ", err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		moduleMass, err := strconv.Atoi(moduleMassString)
		if err != nil {
			log.Println("mass is not a number", err)
		} else {
			fmt.Println(scanner.Text()) // Println will add back the final '\n'
			modules = append(modules, NewModule(moduleMass))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
