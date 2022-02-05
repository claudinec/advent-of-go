package advent2021

import (
	"log"
	"os"
)

// Sonar Sweep: How many measurements are larger than the previous measurement?
func Day01() {
	data, err := os.ReadFile("input/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// do stuff
}

func Main() {
	os.Chdir("advent2021")
	Day01()
}
