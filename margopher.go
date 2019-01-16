package margopher

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

type margopher struct {
	states map[[2]string][]string
}

func New() *margopher {
	return &margopher{}
}

func (m *margopher) ReadText(text string) string {
	m.parse(text)

	return m.generate()
}

func (m *margopher) ReadFile(filePath string) string {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Read data from the file
	text, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	m.parse(string(text))

	return m.generate()
}

func (m *margopher) ReadURL(URL string) string {
	// Open web page
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}

	// Search for <p></p> under <article></article> tags
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		text := s.Find("p").Text()
		m.parse(text)
	})

	return m.generate()
}

func (m *margopher) StateDictionary() map[[2]string][]string {
	return m.states
}

// Parse input text into states map
func (m *margopher) parse(text string) {
	// Initialize margopher.states map
	m.states = make(map[[2]string][]string)

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

// Generate margopher senetence based on a given length
func (m *margopher) generate() string {
	var sentence bytes.Buffer

	// Initialize prefix with a random key
	prefix := m.getRandomPrefix([2]string{"", ""})
	sentence.WriteString(strings.Join(prefix[:], " ") + " ")

	for {
		suffix := getRandomWord(m.states[prefix])
		sentence.WriteString(suffix + " ")

		// Break the loop if suffix ends in "." and senetenceLength is enough
		if isTerminalWord(suffix) {
			break
		}

		prefix = [2]string{prefix[1], suffix}
	}

	return sentence.String()
}

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
	if cap(slice) != 0 {
		return slice[rand.Intn(len(slice))]
	} else {
		return ""
	}
}

func isTerminalWord(word string) bool {
	match, _ := regexp.MatchString("(\\.|,|:|;|\\?|!)$", word)
	return match
}
