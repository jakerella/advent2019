package main

import (
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"strings"
	"strconv"
)

func main() {
	h.WriteHeader(4)

	// 785961
	possibles := []int{}
	for i:=271973; i<785961; i++ {
		num := strings.Split(strconv.Itoa(i), "")
		
		holdLastDigit, err := strconv.Atoi(num[0])
		if err != nil { h.OhShit(err) }

		goodPass := false
		for j:=1; j<len(num); j++ {
			digit, err := strconv.Atoi(num[j])
			if err != nil { h.OhShit(err) }
			if digit < holdLastDigit { goodPass = false; break }
			if digit == holdLastDigit { goodPass = true }
			holdLastDigit = digit
		}
		if goodPass { possibles = append(possibles, i) }
	}

	// log.Info(possibles)
	log.Info("Possible passwords (p1): ", len(possibles));
}
