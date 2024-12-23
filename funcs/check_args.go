package Mosdef

import (
	"log"
	"os"
	"strings"
)

func CheckArgs() {
	if len(os.Args) > 2 {
		log.Fatal("ERROR: invalid data format. too many arguments please enter the name of the file only ")
	}
	if len(os.Args) < 2 {
		log.Fatal("ERROR: invalid data format. not enough arguments please enter the name of the file")
	}
	if !strings.HasSuffix(os.Args[1], ".txt") {
		log.Fatal("ERROR: invalid data format. invalid file format please enter a '.txt'file")
	}
}
