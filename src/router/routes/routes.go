package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rota representa todas as rotas da API
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// Configurar coloca todas as rotas dentro do router
func Config(r *mux.Router) *mux.Router {
	routes := rotesUsers

	for _, rota := range routes {
		r.HandleFunc(rota.URI, rota.Function).Methods(rota.Method)
	}
	return r
}
