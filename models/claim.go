package models

import "github.com/golang-jwt/jwt"

type Claim struct {
	Email       string `json:"email"`
	TipoUsuario string `json:"tipoUsuario"`
	IdUser      uint   `json:"iduser"`
	jwt.StandardClaims
}