package gin_middleware

//gin_log_middleware 用于gin最后打印req和rsp的中间件

import (
	"bytes"
	"github.com/FengZhg/goLogin"
	"github.com/FengZhg/go_tools/protocol_go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

type (
	outputFunc     func(*gin.Context, *protocol_go.SingleLogInfo)
	getLogTypeFunc func(*gin.Context) int32
	enrichFunc     func(ctx *gin.Context) *protocol_go.SingleLogInfo
)

//用于打印请求相应情况中间件
type requestLog struct {
	outputCallbacks    []outputFunc   // 输出日志信息的回调函数（默认已有log.info()）
	getLogTypeCallback getLogTypeFunc // 获取日志类型的回调 ctx.FullPath()到int32的映射
	enrichCallback     enrichFunc     // 丰富日志信息的钩子
}

//NewRequestLog 新建请求日志结构体
func NewRequestLog(outputCallbacks []outputFunc, getLogTypeCallback getLogTypeFunc, enrichCallback enrichFunc) *requestLog {
	return &requestLog{
		outputCallbacks:    outputCallbacks,
		getLogTypeCallback: getLogTypeCallback,
		enrichCallback:     enrichCallback,
	}
}

// RequestLogMiddleware 获取后置打印请求和返回体的中间件
func (r *requestLog) RequestLogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r.requestLogMiddleware(ctx)
	}
}

//requestLogMiddleware 打印请求和返回体的中间件
func (r *requestLog) requestLogMiddleware(ctx *gin.Context) {

	// 创建BodyLogWriter
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = blw

	// 等待进一步执行
	ctx.Next()

	// 获取登录态
	loginInfo, _ := goLogin.GetLoginInfo(ctx)
	// 构造日志
	logInfo := &protocol_go.SingleLogInfo{
		LogType:   r.runGetLogType(ctx),
		Id:        loginInfo.GetId(),
		FullPath:  ctx.FullPath(),
		Status:    r.getStatus(ctx),
		Req:       r.getRequestBody(ctx),
		Message:   blw.body.String(),
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		TimeStamp: time.Now().Unix(),
	}

}

//getRequestBody 获取请求参数
func (r *requestLog) getRequestBody(ctx *gin.Context) string {

	// read请求参数
	reqBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Errorf("Pharse Request Json Param Error err = %v", err)
		return "请求参数解析错误"
	}

	return string(reqBytes)
}

//getStatus 获取处理状态
func (r *requestLog) getStatus(ctx *gin.Context) string {
	if ctx.IsAborted() {
		return "失败"
	}
	return "成功"
}

//runGetLogType 获取日志类型
func (r *requestLog) runGetLogType(ctx *gin.Context) int32 {
	if r.getLogTypeCallback != nil {
		return r.getLogTypeCallback(ctx)
	}
	return 0
}
