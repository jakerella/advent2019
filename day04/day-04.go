package main

import (
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"strings"
	"strconv"
)

func main() {
	h.WriteHeader(4)

	// 271973 - 785961
	possibles := []int{}
	for i:=271973; i<785961; i++ {
		chars := strings.Split(strconv.Itoa(i), "")
		num := [6]int{}
		var hold int
		for j:=0; j<len(chars); j++ {
			hold, err := strconv.Atoi(chars[j]);
			num[j] = hold
			if err != nil { h.OhShit(err) }
		}
		if hold < 0 { log.Error("This is dumb") }
		// log.Info("Checking ", num)

		increments := true
		hasPair := false
		for j:=1; j<len(num); j++ {
			if num[j] < num[j-1] { increments = false; break }
		}
		if increments {
			for j:=1; j<len(num); j++ {
				if j > 1 && num[j] == num[j-1] && num[j] == num[j-2] { continue }
				if j < 4 && num[j] == num[j+1] && num[j] == num[j+2] { continue }
				if j == 1 && num[j] == num[j-1] && num[j] != num[j+1] { hasPair = true }
				if j > 1 && j < 5 && num[j] == num[j-1] && num[j] != num[j+1] { hasPair = true }
				if j == 5 && num[j] == num[j-1] { hasPair = true }
			}
		}
		if increments && hasPair { possibles = append(possibles, i) }
	}

	// log.Info(possibles)
	log.Info("Possible passwords (p2): ", len(possibles));
}
