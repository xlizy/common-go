package zlog

import (
	"github.com/xlizy/common-go/const/threadlocal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Log *zap.Logger
var SLog *zap.SugaredLogger

type Logger struct {
	*zap.SugaredLogger
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
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeTime = timeEncoder
	encoderConfig.EncodeCaller = customCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
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

func appendTraceId(template string) string {
	traceId := "<nil>"
	if threadlocal.TraceId.Get() != nil {
		v := threadlocal.TraceId.Get()
		if v != nil {
			traceId = v.(string)
		}
	}
	template = "traceId:" + traceId + "\t" + template
	return template
}

func Info(template string, args ...interface{}) {
	SLog.Infof(appendTraceId(template), args...)
}
