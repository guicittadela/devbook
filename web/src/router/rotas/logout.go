package rotas

import (
	"net/http"
	"web/src/controllers"
)

var rotaLogout = Rota{
	URI:                "/logout",
	Metodo:             http.MethodGet,
	Funcao:             controllers.Logout,
	RequerAutenticacao: true,
}
