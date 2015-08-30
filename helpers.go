package margopher

import (
	"math/rand"
	"regexp"
)

// Return a random element from a given string slice
func getRandomWord(slice []string) string {
	if !(cap(slice) == 0) {
		return slice[rand.Intn(len(slice))]
	} else {
		return ""
	}
}

// Confirm that a string word ends in '.'
func isTerminalWord(word string) bool {
	match, _ := regexp.MatchString("(\\.)$", word)
	return match
}
