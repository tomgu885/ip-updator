package logger

import "go.uber.org/zap"

//var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	logger := zap.Must(zap.NewDevelopment()).WithOptions(zap.AddCallerSkip(1))
	sugar = logger.Sugar()

}

func Sync() {
	_ = sugar.Sync()
}

func Info(format string) {
	sugar.Info(format)
}

func Infof(format string, args ...interface{}) {
	sugar.Infof(format, args...)
}

func Infow(format string, kwargs ...any) {
	sugar.Infow(format, kwargs...)
}

func Error(args ...any) {
	sugar.Error(args...)
}

func Errorf(format string, args ...any) {
	sugar.Errorf(format, args...)
}

func Warn(args ...any) {
	sugar.Warn(args...)
}

func Warnf(format string, args ...any) {
	sugar.Warnf(format, args...)
}
