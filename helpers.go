package margopher

import (
	"math/rand"
	"regexp"
)

// Return a random prefix other than the one in the arguments
func (m *margopher) getRandomPrefix(prefix [2]string) [2]string {
	// By default, Go orders keys randomly for maps
	for key := range m.states {
		if key != prefix {
			prefix = key
			break
		}
	}

	return prefix
}

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
