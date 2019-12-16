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

type Moon struct {
	pos [3]int
	velocity [3]int
}

func main() {
	h.WriteHeader(12)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the input filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	moons := []Moon{}
	scanner := bufio.NewScanner(input)
	moonRe := regexp.MustCompile(`<x=(\-?\d+), y=(\-?\d+), z=(\-?\d+)>`)
    for scanner.Scan() {
		line := scanner.Text()
		moon := moonRe.FindStringSubmatch(line)
		// log.Info("Found moon: ", moon)
		x, err := strconv.Atoi(moon[1])
		if err != nil { h.OhShit(err) }
		y, err := strconv.Atoi(moon[2])
		if err != nil { h.OhShit(err) }
		z, err := strconv.Atoi(moon[3])
		if err != nil { h.OhShit(err) }
		moons = append(moons, Moon{ [3]int{x, y, z}, [3]int{0, 0, 0} })
	}

	startPosX := [][2]int{}
	startPosY := [][2]int{}
	startPosZ := [][2]int{}
	for _,moon := range moons {
		startPosX = append(startPosX, [2]int{moon.pos[0], moon.velocity[0]})
		startPosY = append(startPosY, [2]int{moon.pos[1], moon.velocity[1]})
		startPosZ = append(startPosZ, [2]int{moon.pos[2], moon.velocity[2]})
	}
	// log.Info("Start Positions: ", startPosX, startPosY, startPosZ)

	xMatch := 0
	yMatch := 0
	zMatch := 0
	i := 0
	for {
		i++
		for ia:=0; ia<len(moons); ia++ {
			for ib:=ia+1; ib<len(moons); ib++ {
				moons[ia], moons[ib] = applyGravity(moons[ia], moons[ib])
			}
		}
		for j:=0; j<len(moons); j++ {
			moons[j] = move(moons[j])
		}
		// log.Info("After Time Step ", i+1, ": ", moons)

		if xMatch == 0 && comparePos(startPosX, 0, moons) { xMatch = i }
		if yMatch == 0 && comparePos(startPosY, 1, moons) { yMatch = i }
		if zMatch == 0 && comparePos(startPosZ, 2, moons) { zMatch = i }
		
		if xMatch > 0 && yMatch > 0 && zMatch > 0 {
			break
		}
	}

	log.Info("Matched all 3: ", xMatch, yMatch, zMatch, "... Repeat on step: ", lcm(xMatch, yMatch, zMatch))

	totalEnergy := float64(0)
	for _,moon := range moons {
		potential := math.Abs(float64(moon.pos[0])) + math.Abs(float64(moon.pos[1])) + math.Abs(float64(moon.pos[2]))
		kinetic := math.Abs(float64(moon.velocity[0])) + math.Abs(float64(moon.velocity[1])) + math.Abs(float64(moon.velocity[2]))
		log.Info("Energy: ", potential, kinetic)
		totalEnergy += (potential * kinetic)
	}

	log.Info("That's no moon (p1)... ", totalEnergy)
}

func applyGravity(moonA Moon, moonB Moon) (a Moon, b Moon) {
	if moonA.pos[0] > moonB.pos[0] {
		moonA.velocity[0]--
		moonB.velocity[0]++
	} else if moonA.pos[0] < moonB.pos[0] {
		moonA.velocity[0]++
		moonB.velocity[0]--
	}
	if moonA.pos[1] > moonB.pos[1] {
		moonA.velocity[1]--
		moonB.velocity[1]++
	} else if moonA.pos[1] < moonB.pos[1] {
		moonA.velocity[1]++
		moonB.velocity[1]--
	}
	if moonA.pos[2] > moonB.pos[2] {
		moonA.velocity[2]--
		moonB.velocity[2]++
	} else if moonA.pos[2] < moonB.pos[2] {
		moonA.velocity[2]++
		moonB.velocity[2]--
	}
	return moonA, moonB
}

func move(moon Moon) Moon {
	moon.pos[0] += moon.velocity[0]
	moon.pos[1] += moon.velocity[1]
	moon.pos[2] += moon.velocity[2]
	return moon
}

func comparePos(start [][2]int, dimension int, moons []Moon) bool {
	for i,pos := range start {
		if pos[0] != moons[i].pos[dimension] || pos[1] != moons[i].velocity[dimension] {
			return false
		}
	}
	return true
}


// These two were ruthlessly stolen from https://play.golang.org/p/SmzvkDjYlb

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
