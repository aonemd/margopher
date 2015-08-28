package margopher

import (
	"bytes"
	"math/rand"
	"regexp"
	"strings"
)

type margopher struct {
	states map[string][]string
}

// Margopher constructor
func NewMargopher() *margopher {
	return &margopher{states: make(map[string][]string)}
}

// Parse input text into states map
func (m *margopher) ParseText(text string) {
	words := strings.Split(text, " ")

	for i := 0; i < len(words)-1; i++ {
		if _, ok := m.states[words[i]]; ok {
			m.states[words[i]] = append(m.states[words[i]], words[i+1])
		} else {
			slice := []string{}
			slice = append(slice, words[i+1])
			m.states[words[i]] = slice
		}
	}
}

// Extract keys from states map
func (m *margopher) extractKeys() []string {
	keys := make([]string, 0, len(m.states))
	for k := range m.states {
		keys = append(keys, k)
	}

	return keys
}

// Generate a random element from a string array
func getRandomWord(slice []string) string {
	if !(cap(slice) == 0) {
		return slice[rand.Intn(len(slice))]
	} else {
		return ""
	}
}

// Confirm that a string ends in '.'
func isTerminalWord(word string) bool {
	match, _ := regexp.MatchString("(\\.)$", word)
	return match
}

// Generate margopher senetence based on given length
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

		// Break the loop if the suffix ends in "." and senetenceLength is enough
		if isTerminalWord(suffix) && i > sentenceLength {
			break
		}

		prefix = suffix
	}

	return sentence.String()
}
