package zlog

import (
	"github.com/xlizy/common-go/const/threadlocal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var Log *zap.Logger
var SLog *zap.SugaredLogger

var space map[int]string

type Logger struct {
	*zap.SugaredLogger
}

func init() {
	space = make(map[int]string, 45)
	for i := 0; i < 45; i++ {
		v := ""
		for j := 0; j < i; j++ {
			v += " "
		}
		space[i] = v
	}
}

func InitLogger(path string) {
	//获取编码器
	encoder := getEncoder()
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(getWriterSyncer(path), zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, multiWriteSyncer, zap.InfoLevel)
	//生成Logger
	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	SLog = Log.Sugar()
	defer Log.Sync()
}

// core 三个参数之  Encoder 编码
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = timeEncoder
	encoderConfig.EncodeCaller = customCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + appendTraceId() + "]")
	path := caller.TrimmedPath()
	formatLen := 35
	if len(path) < formatLen {
		path += space[formatLen-len(path)]
	}
	enc.AppendString("[" + path + "]")
}

func getWriterSyncer(path string) zapcore.WriteSyncer {
	//file, _ := os.Create("/Users/xlizy/XLIZY/workspace-go/logs/log.log")
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
	SLog.Debugf(replace(template), args...)
}

func Info(template string, args ...interface{}) {
	SLog.Infof(replace(template), args...)
}

func Warn(template string, args ...interface{}) {
	SLog.Warnf(replace(template), args...)
}

func Error(template string, args ...interface{}) {
	SLog.Errorf(replace(template), args...)
}

func replace(template string) string {
	return strings.Replace(template, "{}", "%v", -1)
}
