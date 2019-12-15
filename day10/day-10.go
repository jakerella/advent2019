package main

import (
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"os"
	"errors"
	"bufio"
	"strings"
	"math"
	"sort"
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
	var visibleAngles map[float64][2]int
	for y,sectorLine := range asteroidMap {
		for x,sector := range sectorLine {
			if sector == 1 {
				visibleAngles = findVisibles([2]int{x,y}, asteroidMap)
				if len(visibleAngles) > holdMaxVisible {
					holdMaxVisible = len(visibleAngles)
					holdPos = [2]int{x,y}
				}
			}
		}
	}

	log.Info("Max visible (p1): ", holdMaxVisible, " at ", holdPos)
	
	visibleAngles = findVisibles(holdPos, asteroidMap)

	numVaporized := 0
	var twoHundredth [2]int
	for asteroidCount > 1 {
		keys := make([]float64, 0, len(visibleAngles))
		for k,_ := range visibleAngles {
			keys = append(keys, k)
		}
		sort.Float64s(keys)
		// I'm off by 90 degrees... no idea why, this finds "up" (0 deg) and shifts the keys
		var up int
		for i,angle := range keys {
			if visibleAngles[angle][0] == holdPos[0] && visibleAngles[angle][1] < holdPos[1] {
				up = i
				break
			}
		}
		temp := keys[up:]
		keys = append(temp, keys[0:up]...)
		// log.Info(keys)

		for _,angle := range keys {
			numVaporized++
			asteroidCount--
			log.Info("Vaporizing asteroid ", numVaporized, " at ", visibleAngles[angle], " with angle ", angle)
			asteroidMap[visibleAngles[angle][1]][visibleAngles[angle][0]] = 0
			if numVaporized == 200 {
				twoHundredth = visibleAngles[angle]
			}
		}

		log.Info("Finding visible asteroids again")
		visibleAngles = findVisibles(holdPos, asteroidMap)
	}
	log.Info("Two Hundredth kill (p2): ", twoHundredth)
}


func findVisibles(pos [2]int, asteroidMap [][]int) map[float64][2]int {
	// log.Info("Finding visible asteroids from point ", pos)
	angles := make(map[float64][2]int)
	for y,sectorLine := range asteroidMap {
		for x,sector := range sectorLine {
			if sector == 1 {
				if pos[0] == x && pos[1] == y { continue }
				angle := getAngle(pos, [2]int{x,y})
				if _,ok := angles[angle]; !ok {
					angles[angle] = [2]int{x,y}
				}	
			}
		}
	}
	return angles
}

func getAngle(source [2]int, target [2]int) float64 {
	angle := math.Atan2(float64(target[1] - source[1]), float64(target[0] - source[0])) * 180 / math.Pi
	if angle < 0 { angle += 360 }
	return angle
}
