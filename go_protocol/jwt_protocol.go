package go_protocol

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtStatus struct {
	LoginStatus
	jwt.RegisteredClaims
}

func (j *JwtStatus) GetLoginStatus() *LoginStatus {
	if j == nil {
		return nil
	}
	return &j.LoginStatus
}

func (j *JwtStatus) GetRegisteredClaims() *jwt.RegisteredClaims {
	if j == nil {
		return nil
	}
	return &j.RegisteredClaims
}
