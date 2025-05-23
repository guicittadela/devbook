package rotas

import (
	"net/http"
	"web/src/controllers"
)

var rotasUsuario = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.PaginaCadastroDeUsuario,
		RequerAutenticacao: false,
	},
}
