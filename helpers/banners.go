package helpers

import (
	"fmt"
	"strconv"
)

func WriteHeader(day int) {
	fmt.Printf(
		"*****************************\n" +
		"Advent of Code 2019 - Day " + strconv.Itoa(day) + "\n*****************************\n")
}
