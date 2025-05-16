package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)

	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	publicacao.AutorID = usuarioID
	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	repositorio := repositorios.NovoRepositorioPublicacao(db)
	ID, erro := repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	publicacao.ID = ID
	respostas.JSON(w, http.StatusCreated, publicacao)
}
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
