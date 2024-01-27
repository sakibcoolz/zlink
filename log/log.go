package log

import (
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	logger *zap.Logger
}

func New() *Log {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			log.Println("service down", err.Error())
		}
	}(logger)

	logger.Info("Log initialized")

	return &Log{logger: logger}
}

func (l *Log) Log() *zap.Logger {
	return l.logger
}

func (l *Log) Syncronize() {
	l.logger.Sync()
}

func (l *Log) Fatal(ctx *gin.Context, str string, fields ...zapcore.Field) {
	header := ctx.Request.Header

	method := ctx.Request.Method

	fields = append(fields, getLogDatils()...)

	l.logger.Fatal(str, zap.Any("header", header),
		zap.String("method", method),
		zap.Any("fileds", fields))
}

func (l *Log) Debug(ctx *gin.Context, str string, fields ...zapcore.Field) {
	header := ctx.Request.Header

	method := ctx.Request.Method

	fields = append(fields, getLogDatils()...)

	l.logger.Debug(str, zap.Any("header", header),
		zap.String("method", method),
		zap.Any("fileds", fields))
}

func (l *Log) Error(ctx *gin.Context, str string, fields ...zapcore.Field) {

	header := ctx.Request.Header

	method := ctx.Request.Method

	fields = append(fields, getLogDatils()...)

	l.logger.Error(str, zap.Any("header", header),
		zap.String("method", method),
		zap.Any("fileds", fields))
}

func (l *Log) Info(ctx *gin.Context, str string, fields ...zapcore.Field) {

	header := ctx.Request.Header

	method := ctx.Request.Method

	fields = append(fields, getLogDatils()...)

	l.logger.Info(str, zap.Any("header", header),
		zap.String("method", method),
		zap.Any("fileds", fields))
}

func getLogDatils() []zapcore.Field {
	zapcore := make([]zapcore.Field, 0)
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return zapcore
	}

	fname := runtime.FuncForPC(pc)

	funcname := fname.Name()

	file, line := fname.FileLine(pc)

	zapcore = append(zapcore, zap.String("funcname", funcname), zap.String("file", file), zap.Int("line", line))

	return zapcore
}
