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

var phases = [][]int{}

func main() {
	h.WriteHeader(7)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	program := []int{}
	h.OhShit(json.NewDecoder(input).Decode(&program))

	// generate all sequences of phases
	possiblePhases := []int{0,1,2,3,4}
	phasePermutations(possiblePhases, len(possiblePhases), len(possiblePhases))
	// log.Info(phases)
	
	holdOuput := 0
	phaseSeq := []int{}
	for _, phase := range phases {
		inputSignal := 0
		for _, phaseInput := range phase {
			outputs, err := intcode(append([]int(nil), program...), []int{phaseInput, inputSignal})
			if err != nil { h.OhShit(err) }
			
			// log.Info("Outputs", outputs)
			if outputs[0] > holdOuput {
				holdOuput = outputs[0]
				phaseSeq = phase
			}

			inputSignal = outputs[0]
		}
	}
	log.Info("Largest Output (p1): ", holdOuput, " from phase seq ", phaseSeq)
}


func phasePermutations(seq []int, size int, count int) {
    if size == 1 {
		phases = append(phases, append([]int(nil), seq...))
        // log.Info(seq, append([]int(nil), seq...))
		return
	}

    for i:=0; i<size; i++ {
        phasePermutations(seq, size-1, count);
        // if size is odd, swap first and last element 
        // if size is even, swap i-th and last element 
        if size % 2 == 0 {
			seq[i], seq[size-1] = seq[size-1], seq[i] 
		} else {
            seq[0], seq[size-1] = seq[size-1], seq[0]
		}
	}
}


func intcode(program []int, input []int) (value []int, err error) {
	defer func() {
        if (recover() != nil) {
            err = errors.New("Address out of bounds")
        }
	}()
	
	outputs := []int{}

	// log.Info("Inputs: ", input)

	var address int = 0
	for program[address] != 99 {
		// log.Info(program)

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
			next := input[0]
			input = input[1:]
			program[program[address+1]] = next
			address += 2

		case 4:
			param := 0
			if paramTypes[0] == 0 { param = program[program[address+1]] } else { param = program[address+1] }
			// log.Info("Output instruction: ", param)
			outputs = append(outputs, param)
			address += 2
		
		case 5:
			test, value := 0, 0
			if paramTypes[0] == 0 { test = program[program[address+1]] } else { test = program[address+1] }
			if paramTypes[1] == 0 { value = program[program[address+2]] } else { value = program[address+2] }
			if test != 0 {
				address = value
			} else {
				address += 3
			}

		case 6:
			test, value := 0, 0
			if paramTypes[0] == 0 { test = program[program[address+1]] } else { test = program[address+1] }
			if paramTypes[1] == 0 { value = program[program[address+2]] } else { value = program[address+2] }
			if test == 0 {
				address = value
			} else {
				address += 3
			}

		case 7:
			x, y := 0, 0
			if paramTypes[0] == 0 { x = program[program[address+1]] } else { x = program[address+1] }
			if paramTypes[1] == 0 { y = program[program[address+2]] } else { y = program[address+2] }
			if x < y { program[program[address+3]] = 1 } else { program[program[address+3]] = 0 }
			address += 4
		
		case 8:
			x, y := 0, 0
			if paramTypes[0] == 0 { x = program[program[address+1]] } else { x = program[address+1] }
			if paramTypes[1] == 0 { y = program[program[address+2]] } else { y = program[address+2] }
			if x == y { program[program[address+3]] = 1 } else { program[program[address+3]] = 0 }
			address += 4
		
		default:
			return nil, errors.New("Unknown instruction: " + strconv.Itoa(instruction) + " processing code: " + strconv.Itoa(program[address]) + " - " + code)
		}
	}
	return outputs, nil
}
