package go_protocol

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtStatus = struct {
	LoginStatus
	*jwt.StandardClaims
}
