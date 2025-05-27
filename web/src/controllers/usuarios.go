package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/respostas"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{ErroAPI: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, erro := http.Post(url, "applications/json", bytes.NewBuffer(usuario))
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
