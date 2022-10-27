package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedro-barreto/Trabalho-SD/commons"
	"github.com/pedro-barreto/Trabalho-SD/routes"
)

func main() {

	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor esta rodando na porta: 9000")
	log.Println(server.ListenAndServe())

}
