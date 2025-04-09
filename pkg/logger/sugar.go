package logger

import "go.uber.org/zap"

// DebugS 调试日志，详尽的程序日志
func DebugS(args ...interface{}) {
	Logger.Sugar().Debug(args...)
}

// Debugf 格式化调试日志
func Debugf(template string, args ...interface{}) {
	Logger.Sugar().Debugf(template, args...)
}

// InfoS 告知类日志
func InfoS(args ...interface{}) {
	Logger.Sugar().Info(args...)
}

// Infof 格式化告知类日志
func Infof(template string, args ...interface{}) {
	Logger.Sugar().Infof(template, args...)
}

// WarnS 警告类日志
func WarnS(args ...interface{}) {
	Logger.Sugar().Warn(args...)
}

// Warnf 格式化警告类日志
func Warnf(template string, args ...interface{}) {
	Logger.Sugar().Warnf(template, args...)
}

// ErrorS 错误时记录，不应该中断程序
func ErrorS(args ...interface{}) {
	Logger.Sugar().Error(args...)
}

// Errorf 格式化错误日志
func Errorf(template string, args ...interface{}) {
	Logger.Sugar().Errorf(template, args...)
}

// FatalS 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func FatalS(args ...interface{}) {
	Logger.Sugar().Fatal(args...)
}

// Fatalf 格式化致命错误日志
func Fatalf(template string, args ...interface{}) {
	Logger.Sugar().Fatalf(template, args...)
}

// With 添加结构化字段
func With(args ...interface{}) *zap.SugaredLogger {
	return Logger.Sugar().With(args...)
}
