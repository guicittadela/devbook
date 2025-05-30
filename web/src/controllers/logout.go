package controllers

import (
	"net/http"
	"web/src/cookies"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
