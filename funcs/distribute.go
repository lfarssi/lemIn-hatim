package Mosdef

import (
	"fmt"
	"strings"
)

func DistributeAnts(paths [][]string, numberOfAnts int) [][]int {
	antDistribution := make([][]int, len(paths))
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path)
	}

	// Distribute ants across all paths
	for i := 1; i <= numberOfAnts; i++ {
		minIndex := 0
		minValue := len(antDistribution[0]) + pathLengths[0]

		for j := 1; j < len(paths); j++ {
			currentValue := len(antDistribution[j]) + pathLengths[j]
			if currentValue < minValue {
				minIndex = j
				minValue = currentValue
			}
		}
		antDistribution[minIndex] = append(antDistribution[minIndex], i)
	}
	return antDistribution
}

func SimulateAntMovement(paths [][]string, antDistribution [][]int) (string) {
	var finalResult string
	type AntPosition struct {
		ant  int
		path int
		step int
	}

	var antPositions []AntPosition
	for pathIndex, ants := range antDistribution {
		for _, ant := range ants {
			antPositions = append(antPositions, AntPosition{ant, pathIndex, 0})
		}
	}

	for len(antPositions) > 0 {
		var moves []string
		var newPositions []AntPosition
		usedLinks := make(map[string]bool)

		for _, pos := range antPositions {
			if pos.step < len(paths[pos.path])-1 {
				currentRoom := paths[pos.path][pos.step]
				nextRoom := paths[pos.path][pos.step+1]
				link := currentRoom + "-" + nextRoom
				if !usedLinks[link] {
					moves = append(moves, fmt.Sprintf("L%d-%s", pos.ant, nextRoom))
					newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
					usedLinks[link] = true
				} else {
					newPositions = append(newPositions, pos)
				}
			}
		}
		if len(moves) > 0 {
			finalResult += strings.Join(moves, " ")
			finalResult += "\n"
		}
		antPositions = newPositions
	}
	return finalResult
}
