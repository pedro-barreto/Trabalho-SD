package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedro-barreto/Trabalho-SD/backend/commons"
	"github.com/pedro-barreto/Trabalho-SD/backend/routes"
)

func main() {

	commons.Migrate()

	router := mux.NewRouter()
	routes.SetAlunoRoutes(router)

	server := http.Server{

		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Servidor esta rodando na porta 8080")
	log.Println(server.ListenAndServe())

}
