package main

import (
	"fmt"
	"math/rand"
)

func montyHallNBoxes(numberOfBoxes int, numberOfOpenedBoxes int, strategy bool) bool {
	boxes := make([]string, numberOfBoxes)
	for i := range boxes {
		boxes[i] = "empty"
	}
	boxes[rand.Intn(numberOfBoxes)] = "prize"

	choice := rand.Intn(numberOfBoxes)

	var revealedBoxes []int
	for i := range boxes {
		if i != choice && boxes[i] == "empty" {
			revealedBoxes = append(revealedBoxes, i)
		}
	}
	revealedBoxes = revealedBoxes[:numberOfOpenedBoxes]

	var remainingboxes []int
	for i := range boxes {
		if i != choice && !contains(revealedBoxes, i) {
			remainingboxes = append(remainingboxes, i)
		}
	}

	finalChoice := choice
	if strategy {
		finalChoice = remainingboxes[rand.Intn(len(remainingboxes))]
	}

	return boxes[finalChoice] == "prize"
}

func contains(slice []int, item int) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func runMontyHallProblem(rounds int, numberOfBoxes int, numberOfOpenedBoxes int, strategy bool) float64 {
	var wins int = 0
	var losses int = 0
	for i := 0; i < rounds; i++ {
		if montyHallNBoxes(numberOfBoxes, numberOfOpenedBoxes, strategy) {
			wins++
		} else {
			losses++
		}
	}
	fmt.Println("Wygrane:", wins, "Przegrane:", losses)
	return (float64(wins) / float64(rounds)) * 100
}

func main() {

	rounds := 100
	numberOfBoxes := 50
	numberOfOpenedBoxes := 48

	wynik := runMontyHallProblem(rounds, numberOfBoxes, numberOfOpenedBoxes, true)
	fmt.Printf("Współczynnik wygranej: %.2f%%\n", wynik)
}
