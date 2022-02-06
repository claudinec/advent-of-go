package advent2021

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Sonar Sweep
func Day01() {
	var sonarReadings []int
	var numReadings int = 0
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
		numReadings++
	}

	// How many measurements are larger than the previous measurement?
	for lineCounter := 0; lineCounter < numReadings-1; lineCounter++ {
		var thisVal int = sonarReadings[lineCounter]
		var nextVal int = sonarReadings[lineCounter+1]
		if nextVal > thisVal {
			increaseCounter++
		}
	}
	fmt.Printf("There were %d increased measurements.\n", increaseCounter)

	// Consider sums of a three-measurement sliding window. How many sums
	// are larger than the previous sum?
	var sumIncCounter = 0
	for sumCounter := 0; sumCounter < numReadings-3; sumCounter++ {
		var thisSum int = sonarReadings[sumCounter] + sonarReadings[sumCounter+1] + sonarReadings[sumCounter+2]
		var nextSum int = sonarReadings[sumCounter+1] + sonarReadings[sumCounter+2] + sonarReadings[sumCounter+3]
		if nextSum > thisSum {
			sumIncCounter++
		}
	}
	fmt.Printf("There were %d increased three-measurement sums.\n", sumIncCounter)

}

// TODO Identify the current/most recent day based on contents of input directory
// and run the day's function.
func GetCurrentDay() {
	inputFiles, err := os.ReadDir("input")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(inputFiles))
}

func Main() {
	os.Chdir("advent2021")
	GetCurrentDay()
	//Day01()
}
