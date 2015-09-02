package margopher

import (
	"bytes"
	"strings"
)

type margopher struct {
	states map[[2]string][]string
}

// Margopher constructor
func NewMargopher() *margopher {
	return &margopher{states: make(map[[2]string][]string)}
}

// Generate margopher senetence based on a given length
func (m *margopher) Generate(sentenceLength int) string {

	var sentence bytes.Buffer

	// Initialize prefix with a random key
	prefix := m.getRandomPrefix([2]string{"", ""})
	sentence.WriteString(strings.Join(prefix[:], " ") + " ")

	for i := 1; i < sentenceLength; i++ {
		suffix := getRandomWord(m.states[prefix])
		sentence.WriteString(suffix + " ")

		// Break the loop if suffix ends in "." and senetenceLength is enough
		if isTerminalWord(suffix) && i > sentenceLength {
			break
		}

		prefix = [2]string{prefix[1], suffix}
	}

	return sentence.String()
}
