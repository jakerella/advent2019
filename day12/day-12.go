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
	x int
	y int
	z int
	velocity [3]int
}

func main() {
	h.WriteHeader(10)
	
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
		moons = append(moons, Moon{ x, y, z, [3]int{0,0,0} })
	}

	for i:=0; i<1000; i++ {

		for ia:=0; ia<len(moons); ia++ {
			for ib:=ia+1; ib<len(moons); ib++ {
				moons[ia], moons[ib] = applyGravity(moons[ia], moons[ib])
			}
		}
		for j:=0; j<len(moons); j++ {
			moons[j] = move(moons[j])
		}
		// log.Info("After Time Step ", i+1, ": ", moons)
	}

	totalEnergy := float64(0)
	for _,moon := range moons {
		potential := math.Abs(float64(moon.x)) + math.Abs(float64(moon.y)) + math.Abs(float64(moon.z))
		kinetic := math.Abs(float64(moon.velocity[0])) + math.Abs(float64(moon.velocity[1])) + math.Abs(float64(moon.velocity[2]))
		log.Info("Energy: ", potential, kinetic)
		totalEnergy += (potential * kinetic)
	}

	log.Info("That's no moon... ", totalEnergy)
}

func applyGravity(moonA Moon, moonB Moon) (a Moon, b Moon) {
	if moonA.x > moonB.x {
		moonA.velocity[0]--
		moonB.velocity[0]++
	} else if moonA.x < moonB.x {
		moonA.velocity[0]++
		moonB.velocity[0]--
	}
	if moonA.y > moonB.y {
		moonA.velocity[1]--
		moonB.velocity[1]++
	} else if moonA.y < moonB.y {
		moonA.velocity[1]++
		moonB.velocity[1]--
	}
	if moonA.z > moonB.z {
		moonA.velocity[2]--
		moonB.velocity[2]++
	} else if moonA.z < moonB.z {
		moonA.velocity[2]++
		moonB.velocity[2]--
	}
	return moonA, moonB
}

func move(moon Moon) Moon {
	moon.x += moon.velocity[0]
	moon.y += moon.velocity[1]
	moon.z += moon.velocity[2]
	return moon
}
