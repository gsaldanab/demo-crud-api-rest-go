package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gsaldanab/demo-crud-api-rest-go/commons"
	"github.com/gsaldanab/demo-crud-api-rest-go/routes"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose en el puerto 9000")
	log.Println(server.ListenAndServe())
}
