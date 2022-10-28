package routes

import (
	"github.com/gorilla/mux"
	"github.com/pedro-barreto/Trabalho-SD/controllers"
)

func SetAlunoRoutes(router *mux.Router) {

	subRoute := router.PathPrefix("/aluno").Subrouter()
	subRoute.HandleFunc("/todos", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/cadastrar", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/apagar/{id}", controllers.Delete).Methods("DELETE")
	subRoute.HandleFunc("/editar/{id}", controllers.Edit).Methods("PUT")
	subRoute.HandleFunc("/buscar/{id}", controllers.Get).Methods("GET")

}
