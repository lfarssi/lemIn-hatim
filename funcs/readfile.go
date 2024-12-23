package Mosdef

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile() []string {
	lines := make([]string, 0)
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("ERROR: invalid data format. there is no such file please check the path")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	startex := false
	endex := false
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "##start" {
			startex = true
		}
		if strings.TrimSpace(scanner.Text()) == "##end" {
			endex = true
		}
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}
	if len(lines) == 0 {
		log.Fatal("ERROR: invalid data format. empty file")
	} else if !startex {
		log.Fatal("ERROR: invalid data format.there is no start room")
	} else if !endex {
		log.Fatal("ERROR: invalid data format.there is no end room")
	}
	
	return lines
}
