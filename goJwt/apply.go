package goJwt

import (
	"github.com/FengZhg/go_tools/go_protocol"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

//ApplyToken 申请jwt的token
func (g *goJwt) ApplyToken(uid string) (string, error) {
	// 获取方法函数
	method := jwt.GetSigningMethod(g.alg)
	sig, err := method.Sign(g.buildBaseClaim(uid).String(), g.privateKey)
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
		StandardClaims: jwt.StandardClaims{
			Issuer:    g.typeKey,
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}
}
