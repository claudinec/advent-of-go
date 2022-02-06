package advent2021

import (
    "fmt"
    "log"
    "os"
    "strings"
)

// Sonar Sweep: How many measurements are larger than the previous measurement?
func Day01() {
	data, err := os.ReadFile("input/1.txt")
	if err != nil {
		log.Fatal(err)
	}
    const lines = strings.Split(data, "\n")
    var increaseCounter int = 0
    for lineCounter := 1; lineCounter < len(lines); lineCounter++ {
        if lines[lineCounter] > lines[lineCounter-1] {
            increaseCounter++
        }
    }
    fmt.Println(increaseCounter, "increased measurements")
}

func Main() {
	os.Chdir("advent2021")
	Day01()
}
