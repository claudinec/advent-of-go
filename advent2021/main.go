package advent2021

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Sonar Sweep: How many measurements are larger than the previous measurement?
func Day01() {
	var sonarReadings []int
	var numLines int = 0
	var increaseCounter int = 0

	data, err := os.ReadFile("input/1.txt")
	if err != nil {
		log.Fatal(err)
	}

	lineScanner := bufio.NewScanner(bytes.NewReader(data))
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		lineInt, err := strconv.Atoi(lineScanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sonarReadings = append(sonarReadings, lineInt)
		numLines++
	}

	for lineCounter := 1; lineCounter < numLines; lineCounter++ {
		var thisVal int = sonarReadings[lineCounter]
		var prevVal int = sonarReadings[lineCounter-1]
		if thisVal > prevVal {
			increaseCounter++
		}
	}
	fmt.Printf("There were %d increased measurements.\n", increaseCounter)

}

func Main() {
	os.Chdir("advent2021")
	Day01()
}
