package goJwt

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/FengZhg/go_tools/go_protocol"
	"github.com/FengZhg/go_tools/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

//JwtES JWT封装
type JwtES struct {
	opts       *jwtESOptions
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

type jwtESOptions struct {
	typeKey, alg, priPath, pubPath, tokenHeaderKey string
	validTime                                      time.Duration
}
type JwtESOption interface {
	apply(*jwtESOptions)
}

type funcJwtESOption struct {
	f func(*jwtESOptions)
}

func (f *funcJwtESOption) apply(do *jwtESOptions) {
	if f != nil {
		f.f(do)
	}
}

//newFuncJwtESOption 新构建初始化函数
func newFuncJwtESOption(f func(options *jwtESOptions)) JwtESOption {
	return &funcJwtESOption{
		f: f,
	}
}

//WithTypeKey 自定义typeKey
func WithTypeKey(typeKey string) JwtESOption {
	return newFuncJwtESOption(func(options *jwtESOptions) {
		options.typeKey = typeKey
	})
}

//WithPriPath 自定义私钥路径
func WithPriPath(priPath string) JwtESOption {
	return newFuncJwtESOption(func(options *jwtESOptions) {
		options.priPath = priPath
	})
}

//WithPubPath 自定义公钥路径
func WithPubPath(pubPath string) JwtESOption {
	return newFuncJwtESOption(func(options *jwtESOptions) {
		options.pubPath = pubPath
	})
}

//WithAlg 自定义ECDSA类型
func WithAlg(alg string) JwtESOption {
	return newFuncJwtESOption(func(options *jwtESOptions) {
		options.alg = alg
	})
}

//WithValidTime 自定义有效时间
func WithValidTime(validTime time.Duration) JwtESOption {
	return newFuncJwtESOption(func(options *jwtESOptions) {
		options.validTime = validTime
	})
}

//WithTokenHeaderKey 自定义TokenHeaderKey
func WithTokenHeaderKey(tokenHeaderKey string) JwtESOption {
	return newFuncJwtESOption(func(options *jwtESOptions) {
		options.tokenHeaderKey = tokenHeaderKey
	})
}

//initES 初始化goJwtES
func initES(options *jwtESOptions) *JwtES {
	// 读取文件私钥
	priStr, err := ioutil.ReadFile(options.priPath)
	if err != nil {
		panic(err)
	}
	// 读取文件公钥
	pubStr, err := ioutil.ReadFile(options.pubPath)
	if err != nil {
		panic(err)
	}
	// 构建私钥
	priKey, err := jwt.ParseECPrivateKeyFromPEM(priStr)
	if err != nil {
		panic(err)
	}
	// 构建公钥
	pubKey, err := jwt.ParseECPublicKeyFromPEM(pubStr)
	if err != nil {
		panic(err)
	}
	return &JwtES{
		privateKey: priKey,
		publicKey:  pubKey,
		opts:       options,
	}
}

//默认private key path : ./config/private.ec.key
//默认public key path  : ./config/public.pem
//默认type key         ： default
//initJwtES 获取默认ES选项
func initJwtES(opts ...JwtESOption) *JwtES {
	// 默认结构体
	j := &jwtESOptions{
		priPath:        defaultPrivatePath,
		pubPath:        defaultPublicPath,
		typeKey:        defaultTypeKey,
		alg:            jwt.SigningMethodES512.Alg(),
		tokenHeaderKey: defaultTokenHeaderKey,
		validTime:      45 * time.Minute,
	}

	// 丰富参数
	for _, opt := range opts {
		opt.apply(j)
	}

	// 初始化ES goJwt
	return initES(j)
}

//NewES512 初始化ES512 JWT
func NewES512(opts ...JwtESOption) *JwtES {
	opts = append(opts, WithAlg(jwt.SigningMethodES512.Alg()))
	return initJwtES(opts...)
}

//-----------------------------------------------------//
// 申请jwt ES token相关
//-----------------------------------------------------//

//ApplyToken 申请jwt的token
func (g *JwtES) ApplyToken(uid string, name string) (string, error) {
	return jwt.NewWithClaims(jwt.GetSigningMethod(g.opts.alg), g.buildBaseClaim(uid, name)).SignedString(g.privateKey)
}

//buildBaseClaim 构造登录态
func (g *JwtES) buildBaseClaim(uid string, name string) *go_protocol.JwtStatus {
	return &go_protocol.JwtStatus{
		LoginStatus: go_protocol.LoginStatus{
			Name: name,
			Uid:  uid,
			Type: g.opts.typeKey,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    utils.GetExeFileName() + g.opts.typeKey,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(g.opts.validTime)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
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
	return ctx.GetHeader(g.opts.tokenHeaderKey)
}

//authMiddleware jwt身份校验中间件
func (g *JwtES) authMiddleware(ctx *gin.Context) {
	//从请求头部获取jwt身份描述
	claimStr := g.getClaimStrFromHeader(ctx)
	// 进行校验
	jwtToken, err := jwt.ParseWithClaims(claimStr, &go_protocol.JwtStatus{},
		func(token *jwt.Token) (interface{}, error) {
			if !checkSigningMethodType(g.opts.alg, token.Method) {
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
	if !ok || !jwtToken.Valid || jwtStatus.GetType() != g.opts.typeKey {
		log.Errorf("User Jwt Token Error relfect ok:%v\tToken Valid:%v", ok, jwtToken.Valid)
		ctx.Error(go_protocol.LoginInfoError)
		ctx.Abort()
		return
	}
	ctx.Set(contextTokenKey, *jwtStatus)
	ctx.Next()
}
