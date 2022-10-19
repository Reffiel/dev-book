package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Conexão com Mysql
	ConnectionDBString = ""
	//Porta onde a API vai rodar
	Gateway = 0
)

//Load vai inicializar as variáveis de ambiente
func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Gateway, erro = strconv.Atoi(os.Getenv("API_GATEWAY"))
	if erro != nil {
		Gateway = 9000
	}

	ConnectionDBString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTimeTrue&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
}
