package main

import (
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"os"
	"errors"
	"bufio"
	"strings"
	"math"
)

func main() {
	h.WriteHeader(10)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the input filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	asteroidMap := [][]int{}
	asteroidCount := 0
    scanner := bufio.NewScanner(input)
    for scanner.Scan() {
		line := scanner.Text()
		sectors := strings.Split(line, "")
		sectorLine := []int{}
		for _,sector := range sectors {
			object := 0
			if sector == "#" {
				object = 1
				asteroidCount++
			}
			sectorLine = append(sectorLine, object)
		}
		asteroidMap = append(asteroidMap, sectorLine)
	}

	holdMaxVisible := 0
	var holdPos [2]int
	for y,sectorLine := range asteroidMap {
		for x,sector := range sectorLine {
			if sector == 1 {
				visible := findVisibles([2]int{x,y}, asteroidMap)
				if visible > holdMaxVisible {
					holdMaxVisible = visible
					holdPos = [2]int{x,y}
				}
				// log.Info("Asteroid ", x, y, " can see ", visible)
			}
		}
	}

	log.Info("Max visible (p1): ", holdMaxVisible, " at ", holdPos)
}

func findVisibles(pos [2]int, asteroidMap [][]int) int {
	visible := 0
	angles := make(map[float64]bool)
	for y,sectorLine := range asteroidMap {
		for x,sector := range sectorLine {
			if sector == 1 {
				if pos[0] == x && pos[1] == y { continue }
				angle := getAngle(pos, [2]int{x,y})
				if _,ok := angles[angle]; !ok {
					angles[angle] = true
					visible++
				}	
			}
		}
	}
	return visible
}

func getAngle(source [2]int, target [2]int) float64 {
	angle := math.Atan2(float64(target[0] - source[0]), float64(source[1] - target[1])) * 180 / math.Pi
	if angle < 0 {
		angle = angle + 360
	}
	return angle
}
