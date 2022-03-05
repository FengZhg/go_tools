package gin_middleware

//gin_log_middleware 用于gin最后打印req和rsp的中间件

import (
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//Write 重载普通write
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

//WriteString 重载ctx.string()的WriteString
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// RequestLogMiddleware 获取后置打印请求和返回体的中间件
func RequestLogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestLogMiddleware(ctx)
	}
}

//requestLogMiddleware 打印请求和返回体的中间件
func requestLogMiddleware(ctx *gin.Context) {

	// 创建BodyLogWriter
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw

	// 等待进一步执行
	ctx.Next()

	// 执行完毕
	log.Debugf("Req Path:\t%v\tRequest:\t%v\tResponse:\t%v", ctx.FullPath(), getRequestBody(ctx), blw.body.String())
}

//getRequestBody 获取请求参数
func getRequestBody(ctx *gin.Context) string {

	// read请求参数
	reqBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Errorf("Pharse Request Json Param Error err = %v", err)
		return "请求参数解析错误"
	}

	return string(reqBytes)
}
