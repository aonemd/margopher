package margopher

import (
	"log"
	"net/http"
)

func (m *margopher) Api() {
	routes := CreateRoutes(m)

	router := NewRouter(routes)

	log.Fatal(http.ListenAndServe(":7878", router))
}
