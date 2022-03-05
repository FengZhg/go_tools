package gin_middleware

import (
	"github.com/FengZhg/errs"
	"github.com/FengZhg/protocol_go/protocol_go"
	"github.com/gin-gonic/gin"
	"net/http"
)

//ReplyMiddleware 获取处理错误响应中间件
func ReplyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		replyMiddleware(ctx)
	}
}

//生成错误响应的返回结果
func replyMiddleware(ctx *gin.Context) {

	// 等待请求响应结束
	ctx.Next()

	// 判断是否有错误
	if len(ctx.Errors) != 0 {
		err := ctx.Errors[0].Err
		ctx.JSON(http.StatusInternalServerError, &protocol_go.StandardRsp{
			Status: errs.GetErrorCode(err),
			Error:  errs.GetErrorMsg(err),
		})
	}
}
