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
func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into usuarios(nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Nome, user.Nick, user.Email, user.Senha)
	if erro != nil {
		return 0, erro
	}

	lastIDInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInserted), nil
}
