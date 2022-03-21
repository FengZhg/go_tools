package go_protocol

import "github.com/FengZhg/go_tools/errs"

//登录态相关错误
//	错误代码段：10xx
const (
	errorLoginInfoError = 1001
)

var (
	LoginInfoEmptyParamError = errs.NewError(errorLoginInfoError, "登录态错误，请重新登录")
)
