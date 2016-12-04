package margopher

import "strings"

// Parse input text into states map
func (m *margopher) parse(text string) {
	words := strings.Split(text, " ")

	for i := 0; i < len(words)-2; i++ {
		// Initialize prefix with two consecutive words as the key
		prefix := [2]string{words[i], words[i+1]}

		// Assign the third word as value to the prefix
		if _, ok := m.states[prefix]; ok {
			m.states[prefix] = append(m.states[prefix], words[i+2])
		} else {
			m.states[prefix] = []string{words[i+2]}
		}
	}
}
