package margopher

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"os"
)

func (m *margopher) ReadText(text string) string {
	m.parse(text)

	return m.Generate()
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

	return m.Generate()
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

	return m.Generate()
}
