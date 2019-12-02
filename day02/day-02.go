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

	partOne := append([]int(nil), program...)
	partOne[1] = 12
	partOne[2] = 2
	value, err := intcode(partOne)
	if err != nil { log.Warn(err) }
	log.Info("New address 0 (p1): ", value)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			test := append([]int(nil), program...)
			test[1] = i
			test[2] = j
			result, err := intcode(test)
			if err != nil {
				log.Warn(err)
				continue
			}
			if result == 19690720 {
				log.Info("Found target (p2), inputs: ", i, j, " .. 100 * noun + verb: ", (100 * i + j))
				return
			} // else {
			//	log.Info("Wrong result: ", result, " - inputs: ", i, j)
			//}
		}
	}
	log.Info("Did not find target result")
}

func intcode(program []int) (value int, err error) {
	defer func() {
        if (recover() != nil) {
            err = errors.New("Address out of bounds")
        }
	}()
	
	var address int = 0
	for program[address] != 99 {
		if program[address] == 1 {
			program[program[address+3]] = program[program[address+1]] + program[program[address+2]]
		} else if program[address] == 2 {
			program[program[address+3]] = program[program[address+1]] * program[program[address+2]]
		}
		address += 4
	}
	return program[0], nil
}