package main

import (
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"os"
	"errors"
	"encoding/json"
	"strconv"
	"strings"
)

func main() {
	h.WriteHeader(2)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the input filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	program := []int{}
	h.OhShit(json.NewDecoder(input).Decode(&program))

	partOne := append([]int(nil), program...)

	var programInput int = 1

	value, err := intcode(partOne, programInput)
	if err != nil { h.OhShit(err) }

	log.Info("New address 0 (p1): ", value)
}

func intcode(program []int, input int) (value int, err error) {
	defer func() {
        if (recover() != nil) {
            err = errors.New("Address out of bounds")
        }
	}()
	
	var address int = 0
	for program[address] != 99 {
		code := strconv.Itoa(program[address])
		var instruction int
		var paramTypes = [3]int{0,0,0}
		if program[address] < 100 {
			instruction = program[address]
		} else {
			if len(code) == 3 { code = "00" + code }
			if len(code) == 4 { code = "0" + code }

			chars := strings.Split(code, "")
			instruction, err = strconv.Atoi(chars[3] + chars[4])
			if err != nil { h.OhShit(err) }
			paramTypes[0], err = strconv.Atoi(chars[2])
			if err != nil { h.OhShit(err) }
			paramTypes[1], err = strconv.Atoi(chars[1])
			if err != nil { h.OhShit(err) }
			paramTypes[2], err = strconv.Atoi(chars[0])
			if err != nil { h.OhShit(err) }
		}

		switch instruction {
		case 1:
			x, y := 0, 0
			if paramTypes[0] == 0 { x = program[program[address+1]] } else { x = program[address+1] }
			if paramTypes[1] == 0 { y = program[program[address+2]] } else { y = program[address+2] }
			program[program[address+3]] = x + y
			address += 4

		case 2:
			x, y := 0, 0
			if paramTypes[0] == 0 { x = program[program[address+1]] } else { x = program[address+1] }
			if paramTypes[1] == 0 { y = program[program[address+2]] } else { y = program[address+2] }
			program[program[address+3]] = x * y
			address += 4

		case 3:
			program[program[address+1]] = input
			address += 2

		case 4:
			param := 0
			if paramTypes[0] == 0 { param = program[program[address+1]] } else { param = program[address+1] }
			log.Info("Output instruction: ", param)
			address += 2
		default:
			return 0, errors.New("Unknown instruction: " + strconv.Itoa(instruction) + " processing code: " + strconv.Itoa(program[address]) + " - " + code)
		}
	}
	return program[0], nil
}