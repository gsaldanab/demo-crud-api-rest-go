package routes

import (
	"github.com/gorilla/mux"
	"github.com/gsaldanab/demo-crud-api-rest-go/controllers"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("DELETE")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
}
