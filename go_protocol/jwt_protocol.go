package go_protocol

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtStatus struct {
	LoginStatus
	jwt.StandardClaims
}

func (j *JwtStatus) GetLoginStatus() *LoginStatus {
	if j == nil {
		return nil
	}
	return &j.LoginStatus
}

func (j *JwtStatus) GetStandardClaims() *jwt.StandardClaims {
	if j == nil {
		return nil
	}
	return &j.StandardClaims
}
