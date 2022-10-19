package controlers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Ler o request.body e colocar dentro de um struct
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}
	db, erro := banco.Connect()
	if erro != nil {
		log.Fatal(erro)
	}
	repository := repositories.NewUserRepository(db)
	userID, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", userID)))
}
func SearchAllUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users"))
}
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search user"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
