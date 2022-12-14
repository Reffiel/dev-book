package models

import "time"

//Usuario representa um usuário utilizando a rede social
type User struct {
	ID       uint64    `json: "id,omitempty"`
	Nome     string    `json: "nome,omitempty"`
	Nick     string    `json: "nick,omitempty"`
	Email    string    `json: "email,omitempty"`
	Senha    string    `json: "senha,omitempty"`
	CriadoEm time.Time `json: "criadoEm,omitempty"`
}
