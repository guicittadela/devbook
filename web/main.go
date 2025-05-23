package main

import (
	"log"
	"net/http"
	"web/src/router"
	"web/src/utils"
)

func main() {
	utils.CarregarTemplates()

	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":3000", r))
}
