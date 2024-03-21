package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type CustomJsonEncoder struct {
	zapcore.Encoder
}

func NewCustomJSONEncoder(cfg zapcore.EncoderConfig) CustomJsonEncoder {
	c := &CustomJsonEncoder{zapcore.NewJSONEncoder(cfg)}
	return *c
}

func (enc *CustomJsonEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	if fields == nil {
		fields = make([]zapcore.Field, 0)
	}
	fields = append(fields, zap.String("traceId", appendTraceId()))
	return enc.Encoder.EncodeEntry(ent, fields)
}
