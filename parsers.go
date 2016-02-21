package margopher

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Parse input text into states map
func (m *margopher) ParseText(text string) {
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

// Parse text from file and send it to ParseText
func (m *margopher) ParseFile(filePath string) {
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

	// Call ParseText with the text
	m.ParseText(string(text))
}

// Parse text from URL and send it to ParseText
func (m *margopher) ParseURL(URL string) {
	// Open web page
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}

	// Search for <p></p> under <article></article> tags
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		text := s.Find("p").Text()
		// Call ParseText with the text
		m.ParseText(text)
	})
}
