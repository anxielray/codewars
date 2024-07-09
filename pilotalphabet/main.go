package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	fmt.Println(NATO(os.Args[1]))
}

func NATO(s string) (nato string) {
	var natoMap = map[string]string{
		"A": "Alfa",
		"B": "Bravo",
		"C": "Charlie",
		"D": "Delta",
		"E": "Echo",
		"F": "Foxtrot",
		"G": "Golf",
		"H": "Hotel",
		"I": "India",
		"J": "Juliet",
		"K": "Kilo",
		"L": "Lima",
		"M": "Mike",
		"N": "November",
		"O": "Oscar",
		"P": "Papa",
		"Q": "Quebec",
		"R": "Romeo",
		"S": "Sierra",
		"T": "Tango",
		"U": "Uniform",
		"V": "Victor",
		"W": "Whiskey",
		"X": "X-ray",
		"Y": "Yankee",
		"Z": "Zulu",
	}
	for _, v := range strings.ToUpper(s) {
		nato += natoMap[string(v)] + " "
		if string(v) == "!" || string(v) == "." || string(v) == "," || string(v) == "?" {
			nato += string(v)
		}
	}
	return strings.TrimSpace(strings.ReplaceAll(nato, "  ", " "))
}
