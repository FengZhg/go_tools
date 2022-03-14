package gin_middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//Write 重载普通write
func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

//WriteString 重载ctx.string()的WriteString
func (w bodyWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
