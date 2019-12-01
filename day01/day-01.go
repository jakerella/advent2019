package main

import (
	"advent2019/helpers"
	"strconv"
	"os"
	"io/ioutil"
	"encoding/json"
	"math"
)

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
		moduleFuel := fuelForWeight(mod)
		fuelForFuel := fuelForWeight(moduleFuel)
		for fuelForFuel > 0 {
			moduleFuel += fuelForFuel
			fuelForFuel = fuelForWeight(fuelForFuel)
		}
		totalFuel += moduleFuel
		helpers.Log(strconv.Itoa(mod) + " uses " + strconv.Itoa(moduleFuel))
	}

	helpers.Log("Fuel Required: " + strconv.Itoa(totalFuel))
}

func fuelForWeight(amt int) int {
	var div float64 = float64(amt / 3)
	fuel := int(math.Floor(div)) - 2
	return fuel
}
