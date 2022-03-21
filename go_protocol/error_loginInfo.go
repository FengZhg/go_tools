package go_protocol

import "github.com/FengZhg/go_tools/errs"

//登录态相关错误
//	错误代码段：10xx
const (
	errorLoginInfo         = 1001
	errorLoginInfoNotExist = 1002
)

var (
	LoginInfoError         = errs.NewError(errorLoginInfo, "登录态错误或者过期，请重新登录")
	LoginInfoNotExistError = errs.NewError(errorLoginInfoNotExist, "登录态不存在")
)
