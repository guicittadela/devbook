package main

import (
	"log"
	"net/http"
	"web/src/config"
	"web/src/cookies"
	"web/src/router"
	"web/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
