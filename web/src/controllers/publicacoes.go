package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/requisicoes"
	"web/src/respostas"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{ErroAPI: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{ErroAPI: erro.Error()})
		return
	}

	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
	}
	respostas.JSON(w, response.StatusCode, nil)
}
