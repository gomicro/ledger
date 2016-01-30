package ledger

import (
	"fmt"
	"os"
)

var (
	std = New(os.Stdout, DebugLevel)
)

func Debug(args ...interface{}) {
	std.write(DebugLevel, args...)
}

func Debugf(f string, args ...interface{}) {
	std.write(DebugLevel, fmt.Sprintf(f, args...))
}

func Info(args ...interface{}) {
	std.write(InfoLevel, args...)
}

func Infof(f string, args ...interface{}) {
	std.write(InfoLevel, fmt.Sprintf(f, args...))
}

func Warn(args ...interface{}) {
	std.write(WarnLevel, args...)
}

func Warnf(f string, args ...interface{}) {
	std.write(WarnLevel, fmt.Sprintf(f, args...))
}

func Error(args ...interface{}) {
	std.write(ErrorLevel, args...)
}

func Errorf(f string, args ...interface{}) {
	std.write(ErrorLevel, fmt.Sprintf(f, args...))
}

func Fatal(args ...interface{}) {
	std.write(FatalLevel, args...)
}

func Fatalf(f string, args ...interface{}) {
	std.write(FatalLevel, fmt.Sprintf(f, args...))
}
