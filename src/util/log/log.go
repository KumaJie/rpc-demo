package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"runtime"
)

var logger *zap.Logger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger = zap.New(core)
}

func Info(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	callerFields = append(callerFields, fields...)
	logger.Info(message, callerFields...)
}

func Debug(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	callerFields = append(callerFields, fields...)
	logger.Debug(message, callerFields...)
}

func Error(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	callerFields = append(callerFields, fields...)
	logger.Error(message, callerFields...)
}

func Warn(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	callerFields = append(callerFields, fields...)
	logger.Warn(message, callerFields...)
}

func getCallerInfoForLog() (callerFields []zap.Field) {

	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的调用方的函数信息
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名

	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
	return
}
