package main

import (
	"fmt"
	"github.com/AhmedZaleh/margopher/margopher"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	text := "Andrey Andreyevich Markov was born in Ryazan as the son of the secretary of the public forest management of Ryazan, Andrey Grigorevich Markov, and his first wife Nadezhda Petrovna Markova. In the beginning of the 1860s Andrey Grigorevich moved to St. Petersburg to become an asset manager of Ekaterina Aleksandrovna Valvatyeva. "
	margopher := margopher.New()
	margopher.ReadText(text)
	// fmt.Println(margopher.Generate())

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":7878", router))
}

// handlers
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}
