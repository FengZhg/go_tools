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

const (
	// 无分类日志种类
	defaultLogType = 0
)

type (
	outputFunc func(*gin.Context, *protocol_go.SingleLogInfo)
	enrichFunc func(*gin.Context, *protocol_go.SingleLogInfo)
)

//用于打印请求相应情况中间件
type requestLog struct {
	outputCallbacks []outputFunc // 输出日志信息的回调函数（默认已有log.info()）
	enrichHook      enrichFunc   // 丰富日志信息的钩子
}

//NewRequestLog 新建请求日志结构体
func NewRequestLog(outputCallbacks []outputFunc, enrichHook enrichFunc) *requestLog {
	return &requestLog{
		outputCallbacks: outputCallbacks,
		enrichHook:      enrichHook,
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

	// 构建日志信息
	logInfo := r.buildLogInfo(ctx, blw)
	// 处理日志输出回调函数

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

//buildLogInfo 构建一条日志
func (r *requestLog) buildLogInfo(ctx *gin.Context, blw *bodyLogWriter) *protocol_go.SingleLogInfo {

	// 获取登录态
	loginInfo, _ := goLogin.GetLoginInfo(ctx)
	// 构造日志
	logInfo := &protocol_go.SingleLogInfo{
		LogType:   defaultLogType,
		Id:        loginInfo.GetId(),
		FullPath:  ctx.FullPath(),
		Status:    r.getStatus(ctx),
		Req:       r.getRequestBody(ctx),
		Message:   blw.body.String(),
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		TimeStamp: time.Now().Unix(),
	}

	// 跑一下钩子
	if r.enrichHook != nil {
		r.enrichHook(ctx, logInfo)
	}

	return logInfo
}

//stdoutLogWriter 标准输出日志信息
func stdoutLogWriter(ctx *gin.Context, logInfo *protocol_go.SingleLogInfo) {
	log.Info("LoginInfo ID:%v\tFull Path:%v\tReq Body:%v\tRsp:%v", logInfo.GetId(), logInfo.GetFullPath(),
		logInfo.GetReq(), logInfo.GetMessage())
}
