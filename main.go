package main

import (
	"fmt"
	"log"
	"strings"

	Mosdef "Mosdef/funcs"
)

func main() {
	err:=Mosdef.CheckArgs()
	if err != "" {
		fmt.Println(err)
		return
	}
	lines, err := Mosdef.ReadFile()
	if err != "" {
		fmt.Println(err)
		return
	}
	start, end, antsNumber, graph, err := Mosdef.GetRooms(lines)
	if err != "" {
		fmt.Println(err)
		return
	}
	var allPaths [][]string
	if !Mosdef.LinksChecker(graph) {
		allPaths = Mosdef.BreadthFirstSearch(graph, start, end)
	} else {
		allPaths = Mosdef.DepthFirstSearch(graph, start, end)
	}

	filteredPaths := Mosdef.FilterPaths(allPaths)
	antDistribution := Mosdef.DistributeAnts(filteredPaths, antsNumber)
	finalResult, moveCount := Mosdef.SimulateAntMovement(filteredPaths, antDistribution)
	if moveCount < 1 {
		log.Fatal("ERROR: invalid data format. no path found")

	}
	for i, line := range lines {
		if i == 0 || line[0] != '#' || line == "##start" || line == "##end" {
			fmt.Println(strings.TrimSpace(line))
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(finalResult)
}
