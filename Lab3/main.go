package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	tree    = 1
	burning = 2
)

type Wind struct {
	Direction [2]int
	Strength  int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// forest dimensions
	rows, cols := 100, 100

	// probability of a tree being present at a given location
	treeProbability := 0.43

	// forest initialisation
	forest := generateForest(rows, cols, treeProbability)

	// print the original forest
	// fmt.Println("Forest:")
	// for i := range forest {
	// 	fmt.Println(forest[i])
	// }

	// generate a visualisation of the forest
	visualizeForest(forest, "original_forest.png")
	fmt.Println("Visualisation generated: original_forest.png")

	// define the wind (to the right in this case)
	wind := Wind{Direction: [2]int{0, 0}, Strength: 1} // 0, 0 - no wind

	// run the simulation
	burntForest := burnForest(forest, rows, cols, wind)

	// print the burnt forest
	// fmt.Println("Burnt forest:")
	// for i := range burntForest {
	// 	fmt.Println(burntForest[i])
	// }

	// generate a visualisation of the burnt forest
	visualizeForest(burntForest, "burnt_forest.png")
	fmt.Println("Visualisation generated: burnt_forest.png")

	// count the number of burnt trees
	burntTrees := countBurntTrees(burntForest)

	// print the results
	fmt.Println("Burnt trees:", burntTrees)
	fmt.Println("Burnt trees percentage:", float64(burntTrees)/float64(rows*cols)*100)

	// find the optimal tree probability
	numSimulations := 10000

	results := make(map[float64]float64)
	for i := 0; i < numSimulations; i++ {
		treeProbability := rand.Float64()
		forest := generateForest(rows, cols, treeProbability)
		burntForest := burnForest(forest, rows, cols, wind)
		burntTrees := countBurntTrees(burntForest)
		results[treeProbability] = float64(burntTrees)
	}
	fmt.Println("Optimal tree probability:", findOptimalTreeProbability(results)*100)

	// generate a plot
	plotFilename := "burnt_trees.png"
	generatePlot(results, plotFilename)
	fmt.Println("Plot generated:", plotFilename)
}

func burnForest(forest [][]int, rows, cols int, wind Wind) [][]int {
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
			for step := 1; step <= wind.Strength; step++ {
				newX, newY := x+direction.dx*step, y+direction.dy*step

				// calculate the new coordinates considering wind direction
				newX += wind.Direction[0] * step
				newY += wind.Direction[1] * step

				// if the tree is within the forest and is not burning, set it on fire
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols && forest[newX][newY] == tree {
					// put the tree on fire
					forest[newX][newY] = burning
					// add the tree to the queue
					queue = append(queue, []int{newX, newY})
				}
			}
		}
	}
	return forest
}

// generates and returns a forest with trees at random locations
func generateForest(rows, cols int, treeProbability float64) [][]int {
	forest := make([][]int, rows)
	for i := range forest {
		forest[i] = make([]int, cols)
		for j := range forest[i] {
			if rand.Float64() < treeProbability {
				forest[i][j] = tree
			}
		}
	}
	return forest
}

// counts and returns the number of burnt trees in the forest
func countBurntTrees(forest [][]int) int {
	burntTrees := 0
	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] == burning {
				burntTrees++
			}
		}
	}
	return burntTrees
}

// finds and returns the optimal tree density considering both maximizing forest density and minimizing losses
func findOptimalTreeProbability(results map[float64]float64) float64 {
	optimalTreeDensity := 0.0
	bestScore := 0.0
	for treeDensity, loss := range results {
		// Calculate a score that balances maximizing tree density and minimizing loss
		score := treeDensity * (1.0 - loss)
		if score > bestScore {
			bestScore = score
			optimalTreeDensity = treeDensity
		}
	}
	return optimalTreeDensity
}

// generates a plot and saves it to a file
func generatePlot(results map[float64]float64, filename string) error {
	p := plot.New()
	p.Title.Text = "Burnt trees vs tree probability"
	p.X.Label.Text = "Percentage of forestation"
	p.Y.Label.Text = "Ratio of burnt trees to area"

	pts := make(plotter.XYs, len(results))
	i := 0
	for treeProbability, burntRatio := range results {
		pts[i].X = treeProbability
		pts[i].Y = burntRatio
		i++
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	p.Add(s)
	p.Save(8*vg.Inch, 8*vg.Inch, filename)

	return nil
}

// generates a visualization of the forest and saves it to a file
func visualizeForest(forest [][]int, filename string) {
	dc := gg.NewContext(len(forest[0])*10, len(forest)*10)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	for i, row := range forest {
		for j, cell := range row {
			if cell == tree {
				dc.DrawRectangle(float64(j*10), float64(i*10), 10, 10)
				dc.SetRGB(0.25, 0.5, 0.15) // Green for trees
				dc.Fill()
			} else if cell == burning {
				dc.DrawRectangle(float64(j*10), float64(i*10), 10, 10)
				dc.SetRGB(1, 0, 0) // Red for burning trees
				dc.Fill()
			}
		}
	}
	dc.SavePNG(filename)
}
