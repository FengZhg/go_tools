package gin_middleware

//gin_log_middleware 用于gin最后打印req和rsp的中间件

import (
	"bytes"
	"github.com/FengZhg/go_tools/goJwt"
	"github.com/FengZhg/go_tools/go_protocol"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"sync"
	"time"
)

const (
	// 无分类日志种类
	defaultLogType = 0
)

type (
	OutputFunc func(*gin.Context, *go_protocol.SingleLogInfo)
	EnrichFunc func(*gin.Context, *go_protocol.SingleLogInfo)
)

//用于打印请求相应情况中间件
type requestLog struct {
	outputCallbacks []OutputFunc // 输出日志信息的回调函数（默认已有log.info()）
	enrichHook      EnrichFunc   // 丰富日志信息的钩子
}

//NewRequestLog 新建请求日志结构体
func NewRequestLog(outputCallbacks []OutputFunc, enrichHook EnrichFunc) *requestLog {
	//stdCallback 输出时配logrus Formater文件内容，会输出到文件
	outputCallbacks = append(outputCallbacks, stdCallback)
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
	bw := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Writer = bw
	// 等待进一步执行
	ctx.Next()

	// 构建日志信息
	logInfo := r.buildLogInfo(ctx, bw)
	// 简简单单并发执行日志输出回调函数
	r.doOutputCallbacks(ctx, logInfo)
}

//buildLogInfo 构建一条日志
func (r *requestLog) buildLogInfo(ctx *gin.Context, bw *bodyWriter) *go_protocol.SingleLogInfo {
	// 获取登录态
	loginInfo, _ := goJwt.GetLoginInfo(ctx)
	// 构造日志
	logInfo := &go_protocol.SingleLogInfo{
		LogType:   defaultLogType,
		Id:        loginInfo.GetUid(),
		FullPath:  ctx.FullPath(),
		Status:    r.getStatus(ctx),
		Req:       r.getRequestBody(ctx),
		Message:   bw.body.String(),
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		TimeStamp: time.Now().Unix(),
	}

	// 跑一下钩子
	if r.enrichHook != nil {
		r.enrichHook(ctx, logInfo)
	}
	return logInfo
}

//doOutputCallbacks 运行所有地输出回调函数
func (r *requestLog) doOutputCallbacks(ctx *gin.Context, logInfo *go_protocol.SingleLogInfo) {
	// 简简单单并发执行所有输出回调（包括默认的info日志输出回调）
	wg := sync.WaitGroup{}
	for _, outputCallback := range r.outputCallbacks {
		wg.Add(1)
		go func(callback OutputFunc) {
			defer wg.Done()
			callback(ctx, logInfo)
		}(outputCallback)
	}
	wg.Wait()
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

//stdCallback 标准输出日志信息
func stdCallback(ctx *gin.Context, logInfo *go_protocol.SingleLogInfo) {
	log.Infof("LoginInfo ID:%v \tFull Path:%v \tReq Body:%v \tRsp:%v", logInfo.GetId(), logInfo.GetFullPath(),
		logInfo.GetReq(), logInfo.GetMessage())
}
