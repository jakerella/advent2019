package main

import (
	"advent2019/helpers"
	"strconv"
	"os"
	"io/ioutil"
	"encoding/json"
	"math"
)

/*
Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

For example:

For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
For a mass of 1969, the fuel required is 654.
For a mass of 100756, the fuel required is 33583.
The Fuel Counter-Upper needs to know the total fuel requirement. To find it, individually calculate the fuel needed for the mass of each module (your puzzle input), then add together all the fuel values.

What is the sum of the fuel requirements for all of the modules on your spacecraft?
*/

func main() {
	helpers.WriteHeader(1)
	
	if len(os.Args) < 2 {
		helpers.Error("Please provide the input filename")
		return
	}

	input, err := os.Open(os.Args[1])
	if err != nil {
		helpers.Error(err.Error())
		return
	}
	defer input.Close()

	modules := []int{}
	byteValue, _ := ioutil.ReadAll(input)

	json.Unmarshal(byteValue, &modules)

	totalFuel := 0
	for _, mod := range modules {
		// For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
		var div float64 = float64(mod / 3)
		fuel := int(math.Floor(div)) - 2
		totalFuel += fuel
		helpers.Log(strconv.Itoa(mod) + " uses " + strconv.Itoa(fuel))
	}

	helpers.Log("Answer: " + strconv.Itoa(totalFuel))

}
