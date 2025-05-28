package controllers

import (
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/requisicoes"
	"web/src/utils"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

func PaginaCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)

	fmt.Println(response.StatusCode, erro)
	utils.ExecutarTemplate(w, "home.html", nil)
}
