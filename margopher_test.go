package margopher

import "testing"

func TestGetRandomWord(t *testing.T) {
	if getRandomWord([]string{"1", "2", "3"}) == "" {
		t.Error("getRandomWord: it should return a string element from slice.")
	}
}

func TestIsTerminalWord(t *testing.T) {
	if isTerminalWord("Hey.") == false {
		t.Error("isTerminalWord: it should return true for words ending in `.`")
	}

	if isTerminalWord("Hey,") == false {
		t.Error("isTerminalWord: it should return true for words ending in `,`")
	}

	if isTerminalWord("Hey:") == false {
		t.Error("isTerminalWord: it should return true for words ending in `:`")
	}

	if isTerminalWord("Hey;") == false {
		t.Error("isTerminalWord: it should return true for words ending in `;`")
	}

	if isTerminalWord("Hey?") == false {
		t.Error("isTerminalWord: it should return true for words ending in `?`")
	}

	if isTerminalWord("Hey!") == false {
		t.Error("isTerminalWord: it should return true for words ending in `!`")
	}
}

func TestParse(t *testing.T) {
	m := New()
	m.parse("Cats are nice. Cats love pizza, and Cats hates dogs.")

	if m.states == nil {
		t.Error("ParseText: it should initialize states.")
	}
}
