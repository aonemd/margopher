package margopher

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func CreateRoutes(m *margopher) Routes {
	var routes = Routes{
		Route{
			"readFromText",
			"POST",
			"/readtext",
			m.readFromText,
		},
		Route{
			"readFromFile",
			"POST",
			"/readfile",
			m.readFromFile,
		},
		Route{
			"readFromURL",
			"POST",
			"/readurl",
			m.readFromURL,
		},
	}

	return routes
}
