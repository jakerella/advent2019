package main

import (
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"os"
	"errors"
	"encoding/json"
)

func main() {
	h.WriteHeader(2)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the input filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	program := []int{}
	h.OhShit(json.NewDecoder(input).Decode(&program))

	partOne := program
	partOne[1] = 12
	partOne[2] = 2
	log.Info("New address 0 (p1): ", intcode(partOne))
}

func intcode(program []int) int {
	var address int = 0
	for program[address] != 99 {
		if program[address] == 1 {
			program[program[address+3]] = program[program[address+1]] + program[program[address+2]]
		} else if program[address] == 2 {
			program[program[address+3]] = program[program[address+1]] * program[program[address+2]]
		}
		address += 4
	}
	return program[0]
}