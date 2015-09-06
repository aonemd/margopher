package margopher

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (m *margopher) readFromText(w http.ResponseWriter, r *http.Request) {
	var parameters interface{}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &parameters)
	if err != nil {
		panic(err)
	}

	inputText := parameters.(string)
	m.ReadText(inputText)

	fmt.Fprintln(w, m.Generate())
}

func (m *margopher) readFromFile(w http.ResponseWriter, r *http.Request) {
	var parameters interface{}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &parameters)
	if err != nil {
		panic(err)
	}

	inputText := parameters.(string)
	m.ReadFile(inputText)

	fmt.Fprintln(w, m.Generate())
}

func (m *margopher) readFromURL(w http.ResponseWriter, r *http.Request) {
	var parameters interface{}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &parameters)
	if err != nil {
		panic(err)
	}

	inputText := parameters.(string)
	m.ReadURL(inputText)

	fmt.Fprintln(w, m.Generate())
}
