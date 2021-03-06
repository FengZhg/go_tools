package gin_logrus

import (
	"bytes"
	"fmt"
	"github.com/FengZhg/go_tools/utils"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

//----------------------------------------//
// 常量、变量定义
//----------------------------------------//

// 日志输出颜色
const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

// 日志文件输出路径
var (
	logDirPath  = "log/"
	logFilePath = "default.log"
)

//----------------------------------------//
// 初始化加载log format和日志输出方式
//----------------------------------------//

func init() {
	//展示行号
	log.SetReportCaller(true)
	//使用自定义方式构造日志
	log.SetFormatter(&Formatter{})
	//初始化日志输出文件路径
	logFilePath = utils.GetExeFileName() + ".log"
	//初始化并使用
	initLogMultiWriter()
	//初始化替换回车为空格钩子
	initReplaceEnterHook()
}

// 替换回车的钩子结构体
type replaceEnterHook struct{}

func (r *replaceEnterHook) Levels() []log.Level {
	return log.AllLevels
}
func (r *replaceEnterHook) Fire(entry *log.Entry) error {
	entry.Message = strings.ReplaceAll(entry.Message, "\n", " ")
	return nil
}

//initReplaceEnterHook 初始化替换回车的钩子
func initReplaceEnterHook() {
	log.AddHook(&replaceEnterHook{})
}

//----------------------------------------//
// Logrus format定义
//----------------------------------------//

//Formatter 日志结构体
type Formatter struct{}

//Format 产生日志信息函数
func (m *Formatter) Format(entry *log.Entry) ([]byte, error) {

	var b *bytes.Buffer
	var newLog string

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 处理时间情况
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// 根据日志级别获取需要渲染的日志颜色
	levelColor := m.getLevelColor(entry.Level)

	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] %v%s%v %s:%d %s: %s\n",
			timestamp, levelColor, entry.Level, reset, fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		newLog = fmt.Sprintf("[%s] %v%s%v: %s\n", levelColor, entry.Level, reset, timestamp, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

//getLevelColor 根据日志级别 判断需要渲染的颜色
func (m *Formatter) getLevelColor(level log.Level) string {
	switch level {
	case log.DebugLevel | log.InfoLevel:
		return green
	case log.WarnLevel:
		return yellow
	case log.ErrorLevel:
		return red
	default:
		return white
	}
}

//----------------------------------------//
// 日志输出到屏幕和文件的writer定义
//----------------------------------------//

//getRotateWriter 获取到文件输出的writer
func getRotateWriter() io.Writer {
	//初始画rotateLog Writer
	rotateLog, err := rotateLogs.New(
		logDirPath+logFilePath+".%Y%m%d",
		rotateLogs.WithLinkName(logDirPath+logFilePath),
		rotateLogs.WithMaxAge(time.Hour*7*24),
		rotateLogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		log.Errorf("Init Rotatelog Error err = %v", err)
		return nil
	}
	return rotateLog
}

//initLogMultiWriter 初始化日志的多重输出
func initLogMultiWriter() {
	// 构造writer
	writers, rotateLog := []io.Writer{os.Stderr}, getRotateWriter()
	if rotateLog != nil {
		writers = append(writers, rotateLog)
	}

	// 构造日志MultiWriter
	multiWriter := io.MultiWriter(writers...)

	//修改gin的默认输出writer
	gin.DefaultWriter = multiWriter

	// 修改logrus默认输出writer
	log.SetOutput(multiWriter)
}
