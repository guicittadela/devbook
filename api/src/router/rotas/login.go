package rotas

import (
	"api/src/controllers"
	"net/http"
)

var login = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
