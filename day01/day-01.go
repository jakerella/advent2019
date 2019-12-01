package main

import (
	"errors"
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"strconv"
	"os"
	"encoding/json"
	"math"
)

func main() {
	h.WriteHeader(1)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the input filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	modules := []int{}
	h.OhShit(json.NewDecoder(input).Decode(&modules))

	totalFuel := 0
	part1Fuel := 0
	for _, mod := range modules {
		moduleFuel := fuelForWeight(mod)
		part1Fuel += moduleFuel
		fuelForFuel := fuelForWeight(moduleFuel)
		for fuelForFuel > 0 {
			moduleFuel += fuelForFuel
			fuelForFuel = fuelForWeight(fuelForFuel)
		}
		totalFuel += moduleFuel
		// log.Info(strconv.Itoa(mod) + " uses " + strconv.Itoa(moduleFuel))
	}

	log.Info("Module Fuel Required (p1): " + strconv.Itoa(part1Fuel))
	log.Info("Fuel Required (p2): " + strconv.Itoa(totalFuel))
}

func fuelForWeight(amt int) int {
	var div float64 = float64(amt / 3)
	fuel := int(math.Floor(div)) - 2
	return fuel
}
