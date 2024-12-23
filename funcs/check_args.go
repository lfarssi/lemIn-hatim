package Mosdef

import (
	"os"
	"strings"
)

func CheckArgs() string {
	if len(os.Args) != 2 {
		return "ERROR: invalid data format. please enter the name of the file only "

	}
	if !strings.HasSuffix(os.Args[1], ".txt") {
		return "ERROR: invalid data format. invalid file format please enter a '.txt'file"
	}
	return ""
}
