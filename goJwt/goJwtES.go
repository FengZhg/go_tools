package goJwt

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/FengZhg/go_tools/go_protocol"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

//goJwtES JWT封装
type goJwtES struct {
	typeKey, alg string
	privateKey   *ecdsa.PrivateKey
	publicKey    *ecdsa.PublicKey
}

//-----------------------------------------------------//
// 初始化jwt ES 相关
//-----------------------------------------------------//

//NewES512 初始化ES512 JWT
func NewES512(priPath, pubPath, typeKey string) *goJwtES {
	return NewES(priPath, pubPath, typeKey, jwt.SigningMethodES512.Alg())
}

//NewES 初始化goJwtES
func NewES(priPath, pubPath, typeKey, alg string) *goJwtES {
	// 初始化结构体
	g := &goJwtES{
		typeKey: typeKey,
		alg:     alg,
	}

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
	return g
}

//-----------------------------------------------------//
// 申请jwt ES token相关
//-----------------------------------------------------//

//ApplyToken 申请jwt的token
func (g *goJwtES) ApplyToken(uid string) (string, error) {
	return jwt.NewWithClaims(jwt.GetSigningMethod(g.alg), g.buildBaseClaim(uid)).SignedString(g.privateKey)
}

//buildBaseClaim 构造登录态
func (g *goJwtES) buildBaseClaim(uid string) *go_protocol.JwtStatus {
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

//-----------------------------------------------------//
// jwt ES中间件相关
//-----------------------------------------------------//

const (
	jwtHeaderKey = "token"
)

//AuthMiddleware 获取jwt身份校验中间件
func (g *goJwtES) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g.authMiddleware(ctx)
	}
}

//getClaimStrFromHeader 从请求头部获取jwt身份描述信息
func (g *goJwtES) getClaimStrFromHeader(ctx *gin.Context) string {
	return ctx.GetHeader(jwtHeaderKey)
}

//authMiddleware jwt身份校验中间件
func (g *goJwtES) authMiddleware(ctx *gin.Context) {
	//从请求头部获取jwt身份描述
	claimStr := g.getClaimStrFromHeader(ctx)
	// 进行校验
	jwtToken, err := jwt.Parse(claimStr, func(token *jwt.Token) (interface{}, error) {
		if !checkSigningMethodType(g.alg, token.Method) {
			return nil, fmt.Errorf("signing Method Not Match")
		}
		return g.publicKey, nil
	})
	if err != nil {
		log.Errorf("User Jwt Claim Verify Error err:%v", err)
		ctx.Error(go_protocol.LoginInfoError)
		ctx.Abort()
		return
	}

	// 解析登录态信息
	jwtStatus := jwtToken.Claims.(*go_protocol.JwtStatus)
	ctx.Set(contextTokenKey, jwtStatus)
	ctx.Next()
}
