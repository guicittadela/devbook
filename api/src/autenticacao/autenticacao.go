package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(ID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = ID
	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissoes)
	return token.SignedString([]byte(config.SecretKey))

}
