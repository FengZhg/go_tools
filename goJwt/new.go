package goJwt

//NewES512 初始化ES512 JWT
func NewES512(priPath, pubPath, typeKey string) *goJwt {
	g := &goJwt{
		typeKey:      typeKey,
		initCallback: enrichGoJwtES,
		alg:          ES512Alg,
	}
	g.init(priPath, pubPath)
	return g
}
