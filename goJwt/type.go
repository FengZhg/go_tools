package goJwt

import "github.com/golang-jwt/jwt/v4"

var (
	defaultPrivatePath    = "./config/private.ec.key"
	defaultPublicPath     = "./config/public.pem"
	defaultTypeKey        = "default"
	defaultTokenHeaderKey = "token"
)

//checkSigningMethodType 校验签名函数的类型是否正确
func checkSigningMethodType(alg string, method jwt.SigningMethod) (ok bool) {
	switch alg {
	case jwt.SigningMethodES256.Alg(), jwt.SigningMethodES384.Alg(), jwt.SigningMethodES512.Alg():
		_, ok = method.(*jwt.SigningMethodECDSA)
	case jwt.SigningMethodEdDSA.Alg():
		_, ok = method.(*jwt.SigningMethodEd25519)
	case jwt.SigningMethodHS256.Alg(), jwt.SigningMethodHS384.Alg(), jwt.SigningMethodHS512.Alg():
		_, ok = method.(*jwt.SigningMethodHMAC)
	case jwt.SigningMethodRS256.Alg(), jwt.SigningMethodRS384.Alg(), jwt.SigningMethodRS512.Alg():
		_, ok = method.(*jwt.SigningMethodRSA)
	case jwt.SigningMethodPS256.Alg(), jwt.SigningMethodPS384.Alg(), jwt.SigningMethodRS512.Alg():
		_, ok = method.(*jwt.SigningMethodRSAPSS)
	default:
		ok = false
	}
	return ok
}
