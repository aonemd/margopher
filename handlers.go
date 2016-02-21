package margopher

import (
	"encoding/json"
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

	outputSentence := m.ReadText(inputText)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(outputSentence); err != nil {
		panic(err)
	}
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

	outputSentence := m.ReadFile(inputText)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(outputSentence); err != nil {
		panic(err)
	}
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

	outputSentence := m.ReadURL(inputText)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(outputSentence); err != nil {
		panic(err)
	}
}
