package main

import (
	"banco/routes"
	"net/http"
)

func main() {
	routes.Rotas()
	http.ListenAndServe(":8000", nil)
}
