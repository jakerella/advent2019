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
	outputs []int
	complete bool
	relativeBase int
}

func main() {
	h.WriteHeader(9)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	initState := []int{}
	h.OhShit(json.NewDecoder(input).Decode(&initState))

	program := Program{ append([]int(nil), initState...), 0, [2]int{1, 1}, 0, []int{}, false, 0 }

	programState, err := intcode(program)
	if err != nil { h.OhShit(err) }
	
	log.Info("Output (p1): ", programState.outputs)
}

func intcode(program Program) (p Program, err error) {
	// log.Info("Input: ", input)

	for program.state[program.address] != 99 {
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

		// log.Info("Instruction: ", instruction, paramTypes, program.relativeBase)

		switch instruction {
		case 1:
			values, err := getParamValues(paramTypes, 2, program)
			if err != nil { h.OhShit(err) }
			writeLoc := program.state[program.address + 3]
			if paramTypes[2] == 2 {
				writeLoc = program.relativeBase + writeLoc
			}
			program.state = checkLength(program.state, writeLoc)
			program.state[writeLoc] = values[0] + values[1]
			program.address += 4

		case 2:
			values, err := getParamValues(paramTypes, 2, program)
			if err != nil { h.OhShit(err) }
			writeLoc := program.state[program.address + 3]
			if paramTypes[2] == 2 {
				writeLoc = program.relativeBase + writeLoc
			}
			program.state = checkLength(program.state, writeLoc)
			program.state[writeLoc] = values[0] * values[1]
			program.address += 4

		case 3:
			writeLoc := program.state[program.address + 1]
			if paramTypes[0] == 2 {
				writeLoc = program.relativeBase + writeLoc
			}
			program.state = checkLength(program.state, writeLoc)
			program.state[writeLoc] = program.inputs[program.inputIndex]
			if program.inputIndex == 0 { program.inputIndex++ }
			program.address += 2

		case 4:
			values, err := getParamValues(paramTypes, 1, program)
			if err != nil { h.OhShit(err) }
			program.outputs = append(program.outputs, values[0])
			// log.Info("Output instruction: ", program.outputs)
			program.address += 2
			// return program, nil
		
		case 5:
			values, err := getParamValues(paramTypes, 2, program)
			if err != nil { h.OhShit(err) }
			if values[0] != 0 {
				program.address = values[1]
			} else {
				program.address += 3
			}

		case 6:
			values, err := getParamValues(paramTypes, 2, program)
			if err != nil { h.OhShit(err) }
			if values[0] == 0 {
				program.address = values[1]
			} else {
				program.address += 3
			}

		case 7:
			values, err := getParamValues(paramTypes, 2, program)
			if err != nil { h.OhShit(err) }
			writeLoc := program.state[program.address + 3]
			if paramTypes[2] == 2 {
				writeLoc = program.relativeBase + writeLoc
			}
			program.state = checkLength(program.state, writeLoc)
			if values[0] < values[1] { program.state[writeLoc] = 1 } else { program.state[writeLoc] = 0 }
			program.address += 4
		
		case 8:
			values, err := getParamValues(paramTypes, 2, program)
			if err != nil { h.OhShit(err) }
			writeLoc := program.state[program.address + 3]
			if paramTypes[2] == 2 {
				writeLoc = program.relativeBase + writeLoc
			}
			program.state = checkLength(program.state, writeLoc)
			if values[0] == values[1] { program.state[writeLoc] = 1 } else { program.state[writeLoc] = 0 }
			program.address += 4
		
		case 9:
			value := program.state[program.address + 1]
			if paramTypes[0] == 2 {
				value = program.relativeBase + program.state[program.address + 1]
			}
			program.relativeBase = program.relativeBase + value
			program.address += 2
		
		default:
			return program, errors.New("Unknown instruction: " + strconv.Itoa(instruction) + " processing code: " + strconv.Itoa(program.state[program.address]) + " - " + code)
		}
	}
	program.complete = true
	return program, nil
}

func getParamValues(paramTypes [3]int, paramCount int, program Program) (values []int, err error) {
	for i:=0; i<paramCount; i++ {
		paramType := paramTypes[i]
		value := 0
		paramIndex := 1 + i
		switch paramType {
		case 0:
			if program.state[program.address + paramIndex] < len(program.state) { value = program.state[program.state[program.address + paramIndex]] }
		case 1:
			if (program.address + paramIndex) < len(program.state) { value = program.state[program.address + paramIndex] }
		case 2:
			relIndex := program.state[program.address + paramIndex] + program.relativeBase
			program.state = checkLength(program.state, relIndex)
			value = program.state[relIndex]
		default:
			return []int{}, errors.New("Unknown param type: " + strconv.Itoa(paramType))
		}
		values = append(values, value)
	}
	// log.Info("    Values: ", values)
	return values, nil
}

func checkLength(state []int, lastIndex int) []int {
	if len(state) <= lastIndex {
		addedLength := lastIndex - len(state) + 1
		for i:=0; i<addedLength; i++ {
			state = append(state, 0)
		}
	}
	return state
}
