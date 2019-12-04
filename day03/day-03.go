package main

import (
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"os"
	"errors"
	"bufio"
	"regexp"
	"strconv"
	"math"
)

func main() {
	comRe := regexp.MustCompile(`,`)

	h.WriteHeader(3)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the input filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	wires := []string{}
    scanner := bufio.NewScanner(input)
    for scanner.Scan() {
        wires = append(wires, scanner.Text())
	}
	
	directions := [2][]string{}
	for i, wire := range wires {
		directions[i] = comRe.Split(wire, -1)
		// log.Info("Wire: ", directions[i])
	}

	wireOne := plotWire(directions[0])
	wireTwo := plotWire(directions[1])
	// log.Info(wireOne)
	// log.Info(wireTwo)

	intersections := []string{}
	for _, posOne := range wireOne {
		for _, posTwo := range wireTwo {
			if posOne == posTwo {
				intersections = append(intersections, posOne)
			}
		}
	}
	intersections = intersections[1:] // don't want 0,0
	
	// calculate manhatten distance
	// |x1 - x2| + |y1 - y2|
	holdLeastDist := float64(999999999)
	for _, intersection := range intersections {
		coords := comRe.Split(intersection, -1)
		x, err := strconv.Atoi(coords[0])
		if err != nil { h.OhShit(err) }
		y, err := strconv.Atoi(coords[1])
		if err != nil { h.OhShit(err) }
		dist := math.Abs(float64(x)) + math.Abs(float64(y))
		if dist < holdLeastDist {
			holdLeastDist = dist
		}
	}

	log.Info("Least distant intersection (p1): ", holdLeastDist);
}

func plotWire(directions []string) (coords []string) {
	x, y := 0, 0
	coords = []string{}
	coords = append(coords, "0,0")
	for _, instruction := range directions {
		dir := instruction[0:1]
		length, err := strconv.Atoi(instruction[1:])
		if err != nil { h.OhShit(err) }

		for i:=1; i<length+1; i++ {
			if dir == "U" { y-- }
			if dir == "D" { y++ }
			if dir == "L" { x-- }
			if dir == "R" { x ++ }
			coords = append(coords, strconv.Itoa(x) + "," + strconv.Itoa(y))
		}
	}
	return coords
}
