package main

import (
	"fmt"
	"log"
	"strings"

	Mosdef "Mosdef/funcs"
)

func main() {
	Mosdef.CheckArgs()
	lines := Mosdef.ReadFile()
	start, end, antsNumber, graph := Mosdef.GetRooms(lines)
	var allPaths [][]string
	if Mosdef.LinksChecker(graph) {
		allPaths = Mosdef.BreadthFirstSearch(graph, start, end)
	} else {
		allPaths = Mosdef.DepthFirstSearch(graph, start, end)
	}
	filteredPaths := Mosdef.FilterPaths(allPaths)
	antDistribution := Mosdef.DistributeAnts(filteredPaths, antsNumber)
	finalResult, moveCount := Mosdef.SimulateAntMovement(filteredPaths, antDistribution)

	if moveCount < 1 {
		log.Fatal("Error: Invalid data format")
	}
	for _, line := range lines {
		if len(line) > 2 && line[0] != '#' || line == "##start" || line == "##end" {
			fmt.Println(strings.TrimSpace(line))
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(finalResult)
}
