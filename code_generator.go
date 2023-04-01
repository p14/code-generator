package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Word banks for code combinations
var firstWordBank = []string{
	"PEANUT",
	"EXIT",
	"TUNE",
	"BONFIRE",
	"JAW",
	"MOTHER",
}

var secondWordBank = []string{
	"ASTRONAUT",
	"MONSTERS",
	"SOUR",
	"BLEND",
	"PLANE",
}

var thirdWordBank = []string{
	"COCONUT",
	"SUIT",
	"BLUE",
	"ROCKY",
	"MOON",
	"TREETOP",
}

var fourthWordBank = []string{
	"FIELD",
	"PLANE",
	"CLOUD",
	"NUKE",
	"SHOWER",
	"BUGS",
}

// Function to permeate every possible combination
func getCodeOptions() []string {
	var options []string

	for _, firstWord := range firstWordBank {
		for _, secondWord := range secondWordBank {
			for _, thirdWord := range thirdWordBank {
				for _, fourthWord := range fourthWordBank {
					options = append(options, fmt.Sprintf("%s %s %s %s", firstWord, secondWord, thirdWord, fourthWord))
				}
			}
		}
	}

	return options
}

func main() {
	// Get options and shuffle list
	codeOptions := getCodeOptions()

	rand.Seed(
		time.Now().UnixNano(),
	)

	rand.Shuffle(
		len(codeOptions),
		func(i, j int) { codeOptions[i], codeOptions[j] = codeOptions[j], codeOptions[i] },
	)

	// Log each code option line-by-line
	for _, code := range codeOptions {
		fmt.Println(code)
	}
}
