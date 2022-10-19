package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

//Esta func recebe um banco que é aberto pelo controler, cria uma instância do struc users,
//isso é importante porque dentro deste ficam os métodos de iteração direta com o DB
//Fica isolada, e facilita a alteração do DB
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

//Método criar insere um no banco de dados
func (u users) Create(user models.User) (uint64, error) {
	return 0, nil
}
