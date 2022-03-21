package goJwt

import (
	"github.com/FengZhg/go_tools/go_protocol"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"time"
)

//ApplyToken 申请jwt的token
func (g *goJwt) ApplyToken(uid string) (string, error) {
	// 获取方法函数
	sig, err := jwt.NewWithClaims(jwt.GetSigningMethod(g.alg), g.buildBaseClaim(uid)).SignedString(g.privateKey)
	if err != nil {
		log.Errorf("Apply Token Sign Error err:%v", err)
		return "", err
	}
	return sig, nil
}

//buildBaseClaim 构造登录态
func (g *goJwt) buildBaseClaim(uid string) *go_protocol.JwtStatus {
	return &go_protocol.JwtStatus{
		LoginStatus: go_protocol.LoginStatus{
			Uid:  uid,
			Type: g.typeKey,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    g.typeKey,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
}
