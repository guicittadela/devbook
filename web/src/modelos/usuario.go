package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"web/src/config"
	"web/src/requisicoes"
)

type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:publicacoes`
}

func BuscarPerfilUsuario(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacao := make(chan []Publicacao)

	go BuscarDadoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacao, usuarioID, r)

	var (
		usuario    Usuario
		seguidores []Usuario
		seguindo   []Usuario
		publicacao []Publicacao
	)
	for range 4 {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("erro ao buscar o usuário")
			}
			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("erro ao buscar os seguidores")
			}
			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				return Usuario{}, errors.New("erro quem o usuario está seguindo")
			}
			seguindo = seguindoCarregados

		case PublicacoesCarregadas := <-canalPublicacao:
			if PublicacoesCarregadas == nil {
				return Usuario{}, errors.New("erro ao buscar as publicações")
			}
			publicacao = PublicacoesCarregadas
		}

	}
	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacao
	return usuario, nil
}

func BuscarDadoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}

	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}
	canal <- usuario

}
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguidores
}
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}
	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguindo
}

func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioID)
	response, erro := requisicoes.RequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}
	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}
	canal <- publicacoes
}
