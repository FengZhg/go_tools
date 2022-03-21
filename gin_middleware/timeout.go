package gin_middleware

import (
	"github.com/FengZhg/go_tools/errs"
	"github.com/gin-gonic/gin"
	"time"
)

//TimeoutMiddleware 获取超时控制中间件
func TimeoutMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		timeoutMiddleware(ctx)
	}
}

//timeoutMiddleware 超时控制中间件
func timeoutMiddleware(ctx *gin.Context) {

	// 构造任务结束channel
	requestDoneChannel := make(chan struct{})

	// 启用超时控制
	go func() {
		ctx.Next()
		requestDoneChannel <- struct{}{}
	}()

	select {
	case <-time.After(PostTimeLimit):
		ctx.Abort()
		ctx.Error(errs.NewError(1001, "接口响应超时"))
		return
	case <-requestDoneChannel:
		return
	}
}
