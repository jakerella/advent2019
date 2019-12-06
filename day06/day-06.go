package main

import (
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"os"
	"errors"
	"bufio"
	"regexp"
)

func main() {
	comRe := regexp.MustCompile(`\)`)

	h.WriteHeader(6)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the input filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	orbits :=  make(map[string]string)
    scanner := bufio.NewScanner(input)
    for scanner.Scan() {
		line := scanner.Text()
		orbit := comRe.Split(line, -1)
		orbits[orbit[1]] = orbit[0]
	}

	totalOrbits := 0
	for moon, _ := range orbits {
		count := tallyParents(moon, orbits, 0)
		// log.Info("Count for ", moon, " is ", count)
		totalOrbits += count
	}
	log.Info("Direct and Indirect Orbits (p1): ", totalOrbits)


	myPath := tracePath(orbits["YOU"], orbits)
	santaPath := tracePath(orbits["SAN"], orbits)
	for i, planet := range myPath {
		match := indexOf(santaPath, planet)
		if match > -1 {
			log.Info("Transfers to SAN (p2): ", (i + match))
			break
		}
	}
}

func tallyParents(planet string, orbits map[string]string, lastTally int) int {
	if planet == "COM" { return lastTally }
	lastTally++
	return tallyParents(orbits[planet], orbits, lastTally)
}

func tracePath(start string, orbits map[string]string) []string {
	current := start
	trace := []string{}
	for {
		trace = append(trace, current)
		current = orbits[current]
		if current == "COM" { break }
	}

	return trace
}

func indexOf(set []string, elem string) int {
    for i, val := range set { if elem == val { return i } }
    return -1
}
