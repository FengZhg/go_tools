package goJwt

import (
	"crypto/ecdsa"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

const (
	ES512Alg = "ES512"
)

//goJwt JWT封装
type goJwt struct {
	typeKey, alg string
	privateKey   *ecdsa.PrivateKey
	publicKey    *ecdsa.PublicKey
	initCallback func(string, string, *goJwt)
}

//new 初始化JWT封装
func (g *goJwt) init(priPath, pubPath string) {
	if g == nil {
		panic("Lack Of Go Jwt Init Callback")
	}
	g.initCallback(priPath, pubPath, g)
}

//enrichGoJwtES 构建ES512的JWT
func enrichGoJwtES(priPath, pubPath string, g *goJwt) {
	// 读取文件私钥
	priStr, err := ioutil.ReadFile(priPath)
	if err != nil {
		log.Errorf("Read ES Private Key File Error err:%v", err)
		panic(err)
	}
	// 读取文件公钥
	pubStr, err := ioutil.ReadFile(pubPath)
	if err != nil {
		log.Errorf("Read ES Public Key File Error err:%v", err)
		panic(err)
	}
	// 构建私钥
	priKey, err := jwt.ParseECPrivateKeyFromPEM(priStr)
	if err != nil {
		log.Errorf("Parse ES Private Key Error err:%v", err)
		panic(err)
	}
	// 构建公钥
	pubKey, err := jwt.ParseECPublicKeyFromPEM(pubStr)
	if err != nil {
		log.Errorf("Parse ES Private Key Error err:%v", err)
		panic(err)
	}
	// 构建结构体
	g.privateKey = priKey
	g.publicKey = pubKey
}
