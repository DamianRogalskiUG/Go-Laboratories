package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	empty   = iota
	tree    = 1
	burning = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// forest dimensions
	rows, cols := 20, 20

	// probability of a tree being present at a given location
	treeProbability := 0.5

	// forest initialisation
	forest := make([][]int, rows)
	for i := range forest {
		forest[i] = make([]int, cols)
		for j := range forest[i] {
			if rand.Float64() < treeProbability {
				forest[i][j] = tree
			}
		}
	}
	fmt.Println("Las:", forest)

	// random start of the fire
	startRow, startCol := rand.Intn(rows), rand.Intn(cols)
	forest[startRow][startCol] = burning

	// run the simulation
	burnForest(forest, rows, cols)

}

func burnForest(forest [][]int, rows, cols int) {
	startBurningX, startBurningY := rand.Intn(rows), rand.Intn(cols)

	fmt.Println("Punkt startowy pożaru:", startBurningX, startBurningY)
	fmt.Println("Las:", forest[startBurningX][startBurningY])
	if forest[startBurningX][startBurningY] == tree {
		forest[startBurningX][startBurningY] = burning
	}
	fmt.Println("Las:", forest)

}
