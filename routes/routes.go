package routes

import (
	"banco/controllers"
	"net/http"
)

func Rotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("insert", controllers.Insert)
}
