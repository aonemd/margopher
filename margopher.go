package margopher

import (
	"bytes"
	"math/rand"
	"regexp"
)

type margopher struct {
	states map[string][]string
}

// Margopher constructor
func NewMargopher() *margopher {
	return &margopher{states: make(map[string][]string)}
}

// Extract keys from states map
func (m *margopher) extractKeys() []string {
	keys := make([]string, 0, len(m.states))
	for k := range m.states {
		keys = append(keys, k)
	}

	return keys
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

// Generate margopher senetence based on a given length
func (m *margopher) Generate(sentenceLength int) string {
	// Get all prefixes from states maps
	keys := m.extractKeys()

	var sentence bytes.Buffer

	// Initialize prefix with a random key
	prefix := getRandomWord(keys)
	sentence.WriteString(prefix + " ")

	for i := 1; i < sentenceLength; i++ {
		suffix := getRandomWord(m.states[prefix])
		sentence.WriteString(suffix + " ")

		// Break the loop if suffix ends in "." and senetenceLength is enough
		if isTerminalWord(suffix) && i > sentenceLength {
			break
		}

		prefix = suffix
	}

	return sentence.String()
}
