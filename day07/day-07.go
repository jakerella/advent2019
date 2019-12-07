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

type Program struct {
	state []int
	address int
	inputs [2]int
	inputIndex int
	lastOutput int
	complete bool
}

func main() {
	h.WriteHeader(7)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	program := []int{}
	h.OhShit(json.NewDecoder(input).Decode(&program))

	// generate all sequences of phases
	possiblePhases := []int{5,6,7,8,9}
	phasePermutations(possiblePhases, len(possiblePhases), len(possiblePhases))
	// phases = append(phases, []int{9,8,7,6,5})
	// log.Info(phases)

	holdOuput := 0
	phaseSeq := []int{}
	for _, phase := range phases {
		programs := []Program{}
		currentProgram := 0

		// Create each program with correct phase input
		for _, phaseInput := range phase {
			programs = append(programs, Program{ append([]int(nil), program...), 0, [2]int{phaseInput, 0}, 0, -999, false })
		}

		// i := 0  // infinity check
		for {  // i<40 {
			// i++

			// log.Info("RUN ", currentProgram, ": ", programs[currentProgram].inputs, programs[currentProgram].address, programs[currentProgram].complete)
			programState, err := intcode(programs[currentProgram])
			programs[currentProgram] = programState
			if err != nil { h.OhShit(err) }
			
			// log.Info("END ", currentProgram, ": ", programs[currentProgram].lastOutput, programs[currentProgram].address, programs[currentProgram].complete)
			if programs[currentProgram].complete {
				// log.Info("PROGRAM ", currentProgram, " END: ", programs[currentProgram].lastOutput, programs[currentProgram].complete)
				if programs[currentProgram].lastOutput > holdOuput {
					holdOuput = programs[currentProgram].lastOutput
					phaseSeq = phase
				}
				if currentProgram == len(programs) - 1 { break }
			}

			currentProgram++
			if currentProgram >= len(programs) { currentProgram = 0 }
			programs[currentProgram].inputs[1] = programState.lastOutput

			// kick off program A with its phase input and signal (0)
			// when program A hits an output, save program and address
			// use output of A as input signal to kick off B
			// repeat for C, D, E
			// keep track of current program index
			// as programs output, send input to next program index and resume until output
			// as programs terminate, stop sending input
		}
	}
	log.Info("Largest Output (p2): ", holdOuput, " from phase seq ", phaseSeq)
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


func intcode(program Program) (p Program, err error) {
	defer func() {
        if (recover() != nil) {
            err = errors.New("Address out of bounds")
        }
	}()

	// log.Info("Input: ", input)

	for program.state[program.address] != 99 {
		// log.Trace(program)

		code := strconv.Itoa(program.state[program.address])
		var instruction int
		var paramTypes = [3]int{0,0,0}
		if program.state[program.address] < 100 {
			instruction = program.state[program.address]
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

		// log.Info("Instruction: ", instruction, paramTypes)

		switch instruction {
		case 1:
			x, y := 0, 0
			if paramTypes[0] == 0 { x = program.state[program.state[program.address+1]] } else { x = program.state[program.address+1] }
			if paramTypes[1] == 0 { y = program.state[program.state[program.address+2]] } else { y = program.state[program.address+2] }
			program.state[program.state[program.address+3]] = x + y
			program.address += 4

		case 2:
			x, y := 0, 0
			if paramTypes[0] == 0 { x = program.state[program.state[program.address+1]] } else { x = program.state[program.address+1] }
			if paramTypes[1] == 0 { y = program.state[program.state[program.address+2]] } else { y = program.state[program.address+2] }
			program.state[program.state[program.address+3]] = x * y
			program.address += 4

		case 3:
			program.state[program.state[program.address+1]] = program.inputs[program.inputIndex]
			if program.inputIndex == 0 { program.inputIndex++ }
			program.address += 2

		case 4:
			output := -9999
			if paramTypes[0] == 0 { output = program.state[program.state[program.address+1]] } else { output = program.state[program.address+1] }
			program.lastOutput = output
			// log.Info("Output instruction: ", program.lastOutput)
			program.address += 2
			return program, nil
		
		case 5:
			test, value := 0, 0
			if paramTypes[0] == 0 { test = program.state[program.state[program.address+1]] } else { test = program.state[program.address+1] }
			if paramTypes[1] == 0 { value = program.state[program.state[program.address+2]] } else { value = program.state[program.address+2] }
			if test != 0 {
				program.address = value
			} else {
				program.address += 3
			}

		case 6:
			test, value := 0, 0
			if paramTypes[0] == 0 { test = program.state[program.state[program.address+1]] } else { test = program.state[program.address+1] }
			if paramTypes[1] == 0 { value = program.state[program.state[program.address+2]] } else { value = program.state[program.address+2] }
			if test == 0 {
				program.address = value
			} else {
				program.address += 3
			}

		case 7:
			x, y := 0, 0
			if paramTypes[0] == 0 { x = program.state[program.state[program.address+1]] } else { x = program.state[program.address+1] }
			if paramTypes[1] == 0 { y = program.state[program.state[program.address+2]] } else { y = program.state[program.address+2] }
			if x < y { program.state[program.state[program.address+3]] = 1 } else { program.state[program.state[program.address+3]] = 0 }
			program.address += 4
		
		case 8:
			x, y := 0, 0
			if paramTypes[0] == 0 { x = program.state[program.state[program.address+1]] } else { x = program.state[program.address+1] }
			if paramTypes[1] == 0 { y = program.state[program.state[program.address+2]] } else { y = program.state[program.address+2] }
			if x == y { program.state[program.state[program.address+3]] = 1 } else { program.state[program.state[program.address+3]] = 0 }
			program.address += 4
		
		default:
			return program, errors.New("Unknown instruction: " + strconv.Itoa(instruction) + " processing code: " + strconv.Itoa(program.state[program.address]) + " - " + code)
		}
	}
	program.complete = true
	return program, nil
}
