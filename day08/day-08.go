package main

import (
	h "github.com/jakerella/advent2019/helpers"
	log "github.com/sirupsen/logrus"
	"os"
	"errors"
	"encoding/json"
)

func main() {
	h.WriteHeader(5)
	
	if len(os.Args) < 2 { h.OhShit(errors.New("Please provide the filename")) }

	input, err := os.Open(os.Args[1])
	if err != nil { h.OhShit(err) }
	defer input.Close()

	image := []int{}
	h.OhShit(json.NewDecoder(input).Decode(&image))

	// image = []int{0,2,2,2,1,1,2,2,2,2,1,2,0,0,0,0}
	
	rows := 6
	cols := 25
	layerSize := rows * cols
	layers := [][]int{}

	layerStart := 0
	for layerStart < len(image) {
		layers = append(layers, image[layerStart:(layerStart+layerSize)])
		layerStart += layerSize
	}

	holdFewestZeros := layerSize
	var holdLayer []int
	for _, layer := range layers {
		zeros := 0
		for _, pixel := range layer {
			if pixel == 0 { zeros++ }
		}
		if zeros < holdFewestZeros {
			holdFewestZeros = zeros
			holdLayer = layer
		}
	}

	ones := 0
	twos := 0
	for _, pixel := range holdLayer {
		if pixel == 1 { ones++ }
		if pixel == 2 { twos++ }
	}
	
	log.Info("Least zeros layer checksum (p1): ",(ones * twos))

	visiblePixels := []int{}
	for i:=0; i<layerSize; i++ {
		visiblePixels = append(visiblePixels, 2)
		for _, layer := range layers {
			if layer[i] < 2 {
				visiblePixels[i] = layer[i]
				break
			}
		}
	}

	log.Info("Visible Pixels: ", visiblePixels)
	renderImage(visiblePixels)
}


func renderImage(pixels []int) {
	line := ""
	for _, pixel := range pixels {
		char := " "
		if pixel == 0 { char = " " }
		if pixel == 1 { char = "X" }
		line += char
		if len(line) >= 25 {
			log.Info(line)
			line = ""
		}
	}
}
