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

	// print the original forest
	for i := range forest {
		fmt.Println(forest[i])
	}

	// run the simulation
	burntForest := burnForest(forest, rows, cols)

	// print the burnt forest
	for i := range burntForest {
		fmt.Println(burntForest[i])
	}
}

func burnForest(forest [][]int, rows, cols int) [][]int {

	// random starting point for the fire
	startBurningX, startBurningY := rand.Intn(rows), rand.Intn(cols)

	// if the starting point is a tree, set it on fire
	if forest[startBurningX][startBurningY] == tree {
		forest[startBurningX][startBurningY] = burning
	}

	// queue storing coordinates of the locations to check
	queue := [][]int{{startBurningX, startBurningY}}

	// while there are trees to check, keep checking and spreading the fire
	for len(queue) > 0 {
		// get the coordinates of the tree to check
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		// check all the neighbouring trees
		for _, direction := range []struct{ dx, dy int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}} {
			newX, newY := x+direction.dx, y+direction.dy

			// if the tree is within the forest and is not burning, set it on fire
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols && forest[newX][newY] == tree {
				// put the tree on fire
				forest[newX][newY] = burning
				// add the tree to the queue
				queue = append(queue, []int{newX, newY})
			}
		}
	}
	return forest
}
