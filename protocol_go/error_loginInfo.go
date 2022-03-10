package protocol_go

import "github.com/FengZhg/go_tools/errs"

//登录态相关错误
//	错误代码段：10xx
const (
	errorLoginInfoEmptyParam    = 1001
	errorLoginInfoNotExist      = 1002
	errorLoginInfoIPError       = 1003
	errorLoginInfoTokenError    = 1004
	errorLoginInfoPassStepError = 1005
)

var (
	LoginInfoEmptyParamError = errs.NewError(errorLoginInfoEmptyParam, "登录态参数缺失")
	LoginInfoNotExistError   = errs.NewError(errorLoginInfoNotExist, "登录态不存在")
	LoginInfoTokenError      = errs.NewError(errorLoginInfoTokenError, "登录态Token错误，请重新登录")
	LoginInfoPassStepError   = errs.NewError(errorLoginInfoPassStepError, "登录态校验不通过，请重新登录")
	LoginInfoIPErrorError    = errs.NewError(errorLoginInfoIPError, "登录态IP错误，请重新登录")
)
