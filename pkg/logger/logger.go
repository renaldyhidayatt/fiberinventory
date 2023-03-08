package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	Log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

}

func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	Log.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Log.Error(message, fields...)
}
