package main

import (
	"fmt"
	"math/rand"
	"time"
)

// WORD BANKS FOR CODE COMBINATIONS
var firstWordBank = []string{
	"PEANUT",
	"EXIT",
	"TUNE",
	"BONFIRE",
	"JAW",
	"MOTHER",
	// "FIELD",
	// "PEEP",
}

var secondWordBank = []string{
	"ASTRONAUT",
	"MONSTERS",
	"SOUR",
	"BLEND",
	"PLANE",
	// "PLANTAIN",
	// "NUKE",
	// "PEANUT",
}

var thirdWordBank = []string{
	"COCONUT",
	"SUIT",
	"BLUE",
	"ROCKY",
	"MOON",
	"TREETOP",
	// "RASPBERRY",
	// "EXIT",
}

var fourthWordBank = []string{
	"FIELD",
	"PLANE",
	"CLOUD",
	"NUKE",
	"SHOWER",
	"BUGS",
	// "BEAN",
	// "PEEK",
	// "COCONUT",
}

// FUNCTION TO PERMEATE EVERY POSSIBLE COMBINATION
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
	// GET OPTIONS AND SHUFFLE LIST
	codeOptions := getCodeOptions()

	rand.Seed(
		time.Now().UnixNano(),
	)

	rand.Shuffle(
		len(codeOptions),
		func(i, j int) { codeOptions[i], codeOptions[j] = codeOptions[j], codeOptions[i] },
	)

	// LOG EACH CODE OPTION LINE BY LINE
	for _, code := range codeOptions {
		fmt.Println(code)
	}
}
