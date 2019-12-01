package helpers

import (
	"fmt"
)

func Log(message string) {
	fmt.Printf(message + "\n")
}

func Error(message string) {
	fmt.Printf("ERROR: " + message + "\n")
}
