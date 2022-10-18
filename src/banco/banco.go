package banco

import "database/sql"
_ "github.com/go-sql-driver/mysql" //Driver

//Conect abre a conexão com o banco de dados e retorna
func Conect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnectionDBString)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil{
		db.Close()
		return nil, erro
	}
	return db, nil
}
