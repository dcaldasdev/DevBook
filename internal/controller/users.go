package controller

import (
	"api/internal/entity"
	"api/internal/infra/database"
	"api/internal/infra/database/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Create creates user and persist in the database
func Create(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user entity.User
	if err = json.Unmarshal(request, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	userRepository := repository.NewUserRepository(db)
	userRepository.Create(user)
}

// FindAll finds all users in the database
func FindAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching for all users"))
}

// FindById finds user by ID in the database
func FindById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searching by id"))
}

// Update updates an existing user in the database by ID
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
}

// Delete deletes an existing user in the database by ID
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário!"))
}
