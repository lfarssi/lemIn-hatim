package Mosdef

import (
	"sort"
)

func FilterPaths(allPaths [][]string) [][]string {
	// Sort paths by length to prioritize longer paths
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) > len(allPaths[j])
	})
	// Store the best combination of non-overlapping paths
	var bestCombination [][]string

	// Try all possible combinations to find the optimal set
	var findBestCombination func(currentCombination [][]string, remainingPaths [][]string)
	findBestCombination = func(currentCombination [][]string, remainingPaths [][]string) {
		// If we found a better combination, update bestCombination
		if len(currentCombination) > len(bestCombination) {
			bestCombination = make([][]string, len(currentCombination))
			copy(bestCombination, currentCombination)
		} else if len(currentCombination) == len(bestCombination) {
			if steps(currentCombination) < steps(bestCombination) {
				bestCombination = make([][]string, len(currentCombination))
				copy(bestCombination, currentCombination)
			}
		}

		// Try adding remaining paths that don't over lap
		for i, path := range remainingPaths {
			if canAddPath(currentCombination, path) {
				// Create a new combination with the current path added
				newCombination := make([][]string, len(currentCombination))
				copy(newCombination, currentCombination)
				newCombination = append(newCombination, path)

				// Create new remaining paths excluding the current path and overlapping paths
				newRemaining := make([][]string, 0)
				for j, remainPath := range remainingPaths {
					if i != j && canAddPath(newCombination, remainPath) {
						newRemaining = append(newRemaining, remainPath)
					}
				}

				// Recursively find combinations
				findBestCombination(newCombination, newRemaining)
			}
		}
	}

	// Start with an empty combination and all paths
	findBestCombination([][]string{}, allPaths)

	return bestCombination
}

// Check if a path can be added to the current combination without room overlap
func canAddPath(combination [][]string, newPath []string) bool {
	// Check rooms from index 1 to len-1 to ignore start and end
	for _, existingPath := range combination {
		for _, room1 := range newPath[1 : len(newPath)-1] {
			for _, room2 := range existingPath[1 : len(existingPath)-1] {
				if room1 == room2 {
					return false
				}
			}
		}
	}
	return true
}

func steps(combination [][]string) int {
	count := 0
	for _, i := range combination {
		for _, j := range i {
			if len(j) != 0 {
				count++
			}
		}
	}
	return count
}
