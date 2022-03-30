package errs

import (
	"fmt"
	"net/http"
)

//Err 带错误代码的错误类型
type Err struct {
	Code   int32  // 错误代码
	ErrMsg string // 错误信息
	error
}

//NewError 构造新的错误类型
func NewError(code int32, msg string) Err {
	return Err{
		Code:   code,
		ErrMsg: msg,
	}
}

func (x *Err) GetErrMsg() string {
	return x.ErrMsg
}

func (x *Err) GetCode() int32 {
	return x.Code
}

func (x Err) Error() string {
	return fmt.Sprintf("Error Code: %v \t Error Msg: %v", x.GetCode(), x.GetErrMsg())
}

func (x Err) String() string {
	return x.Error()
}

//GetErrorCode 获取错误代码
func GetErrorCode(err error) int32 {
	errs, ok := err.(Err)
	if !ok {
		return http.StatusInternalServerError
	}
	return errs.GetCode()
}

//GetErrorMsg 获取错误信息
func GetErrorMsg(err error) string {
	errs, ok := err.(Err)
	if !ok {
		return err.Error()
	}
	return errs.GetErrMsg()
}
