package zlog

import (
	"fmt"
	"github.com/xlizy/common-go/base/const/threadlocal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var _log *zap.Logger
var _space map[int]string

type Logger struct {
	*zap.SugaredLogger
}

func InitLogger(path string) {
	_space = make(map[int]string, 45)
	for i := 0; i < 45; i++ {
		v := ""
		for j := 0; j < i; j++ {
			v += " "
		}
		_space[i] = v
	}
	//获取编码器
	encoder := getJsonEncoder()
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(getWriterSyncer(path), zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, multiWriteSyncer, zap.InfoLevel)
	//生成Logger
	_log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	defer _log.Sync()
}

// core 三个参数之  Encoder 编码
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder
	encoderConfig.EncodeCaller = customCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getJsonEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + appendTraceId() + "]")
	path := caller.TrimmedPath()
	formatLen := 35
	if len(path) < formatLen {
		path += _space[formatLen-len(path)]
	}
	enc.AppendString("[" + path + "]")
}

func getWriterSyncer(path string) zapcore.WriteSyncer {
	lumberWriteSyncer := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    30, // megabytes
		MaxBackups: 1000,
		MaxAge:     2000,  // days
		Compress:   false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	}
	return zapcore.AddSync(lumberWriteSyncer)
}

func appendTraceId() string {
	traceId := threadlocal.GetTraceId()
	if traceId == "<nil>" {
		traceId = "00000000-0000-0000-0000-000000000000"
	}
	return traceId
}

func Debug(template string, args ...interface{}) {
	_log.With(zap.String("traceId", appendTraceId())).Debug(fmt.Sprintf(replace(template), args...))
}

func Info(template string, args ...interface{}) {
	_log.With(zap.String("traceId", appendTraceId())).Info(fmt.Sprintf(replace(template), args...))
}

func Warn(template string, args ...interface{}) {
	_log.With(zap.String("traceId", appendTraceId())).Warn(fmt.Sprintf(replace(template), args...))
}

func Error(template string, args ...interface{}) {
	_log.With(zap.String("traceId", appendTraceId())).Error(fmt.Sprintf(replace(template), args...))
}

func replaceOld(template string) string {
	return strings.Replace(template, "{}", "%v", -1)
}

func replace(template string) string {
	return strings.Replace(template, "{}", "%s", -1)
}
