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

//JwtES JWT封装
type JwtES struct {
	TypeKey, Alg, PriPath, PubPath string
	privateKey                     *ecdsa.PrivateKey
	publicKey                      *ecdsa.PublicKey
}

type JwtESOpt func(g *JwtES)

const (
	jwtHeaderKey = "token"
)

//-----------------------------------------------------//
// 初始化jwt ES 相关
//-----------------------------------------------------//

func newESDefault() *JwtES {
	// 默认结构体
	return &JwtES{
		PriPath: privatePath,
		PubPath: publicPath,
		TypeKey: defaultTypeKey,
	}
}

//NewES512 初始化ES512 JWT
//默认private key path : ./config/private.ec.key
//默认public key path  : ./config/public.pem
//默认type key         ： default
func NewES512(opts ...JwtESOpt) *JwtES {
	// 初始化默认GoJwt
	g := newESDefault()
	g.Alg = jwt.SigningMethodES512.Alg()
	// 调用opts函数
	for _, opt := range opts {
		opt(g)
	}
	// 读取key
	g.initECKey()
	return g
}

//initECKey 初始化goJwtES
func (g *JwtES) initECKey() *JwtES {
	// 读取文件私钥
	priStr, err := ioutil.ReadFile(g.PriPath)
	if err != nil {
		log.Errorf("Read ES Private Key File Error err:%v", err)
		panic(err)
	}
	// 读取文件公钥
	pubStr, err := ioutil.ReadFile(g.PubPath)
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
func (g *JwtES) ApplyToken(uid string) (string, error) {
	return jwt.NewWithClaims(jwt.GetSigningMethod(g.Alg), g.buildBaseClaim(uid)).SignedString(g.privateKey)
}

//buildBaseClaim 构造登录态
func (g *JwtES) buildBaseClaim(uid string) *go_protocol.JwtStatus {
	return &go_protocol.JwtStatus{
		LoginStatus: go_protocol.LoginStatus{
			Uid:  uid,
			Type: g.TypeKey,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    g.TypeKey,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
}

//-----------------------------------------------------//
// jwt ES中间件相关
//-----------------------------------------------------//

//AuthMiddleware 获取jwt身份校验中间件
func (g *JwtES) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g.authMiddleware(ctx)
	}
}

//getClaimStrFromHeader 从请求头部获取jwt身份描述信息
func (g *JwtES) getClaimStrFromHeader(ctx *gin.Context) string {
	return ctx.GetHeader(jwtHeaderKey)
}

//authMiddleware jwt身份校验中间件
func (g *JwtES) authMiddleware(ctx *gin.Context) {
	//从请求头部获取jwt身份描述
	claimStr := g.getClaimStrFromHeader(ctx)
	// 进行校验
	jwtToken, err := jwt.ParseWithClaims(claimStr, &go_protocol.JwtStatus{},
		func(token *jwt.Token) (interface{}, error) {
			if !checkSigningMethodType(g.Alg, token.Method) {
				return nil, fmt.Errorf("signing Method Not Match")
			}
			return g.publicKey, nil
		},
	)
	if err != nil {
		log.Errorf("User Jwt Claim Verify Error err:%v", err)
		ctx.Error(go_protocol.LoginInfoError)
		ctx.Abort()
		return
	}

	// 解析登录态信息
	jwtStatus, ok := jwtToken.Claims.(*go_protocol.JwtStatus)
	if !ok || !jwtToken.Valid {
		log.Errorf("User Jwt Token Error relfect ok:%v\tToken Valid:%v", ok, jwtToken.Valid)
		ctx.Error(go_protocol.LoginInfoError)
		ctx.Abort()
		return
	}
	ctx.Set(contextTokenKey, *jwtStatus)
	ctx.Next()
}
