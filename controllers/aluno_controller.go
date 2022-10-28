package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pedro-barreto/Trabalho-SD/commons"
	"github.com/pedro-barreto/Trabalho-SD/models"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	alunos := []models.Aluno{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&alunos)
	json, _ := json.Marshal(alunos)
	commons.SendReponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	aluno := models.Aluno{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()

	db.Find(&aluno, id)

	if aluno.ID > 0 {
		json, _ := json.Marshal(aluno)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func Save(writer http.ResponseWriter, request *http.Request) {
	aluno := models.Aluno{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&aluno)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&aluno).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(aluno)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func Edit(writer http.ResponseWriter, request *http.Request) {
	aluno := models.Aluno{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&aluno, id)

	error := json.NewDecoder(request.Body).Decode(&aluno)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&aluno).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(aluno)

	commons.SendReponse(writer, http.StatusCreated, json)

}

func Delete(writer http.ResponseWriter, request *http.Request) {
	aluno := models.Aluno{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&aluno, id)

	if aluno.ID > 0 {
		db.Delete(aluno)
		commons.SendReponse(writer, http.StatusOK, []byte(``))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
