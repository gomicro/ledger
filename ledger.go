package ledger

import (
	"fmt"
	"io"
)

type Ledger struct {
	Writer    io.Writer
	Threshold Level
}

func (l *Ledger) Write(lvl Level, args ...interface{}) {
	if lvl <= l.Threshold {
		out, args := fmt.Sprintf("[%v]", args[0]), args[1:]
		for _, arg := range args {
			out += fmt.Sprintf(" [%v]", arg)
		}
		fmt.Fprintf(l.Writer, "%s: %v", lvl, out)
	}
}

func (l *Ledger) Debug(args ...interface{}) {
	l.Write(DebugLevel, args...)
}

func (l *Ledger) Debugf(f string, args ...interface{}) {
	l.Write(DebugLevel, fmt.Sprintf(f, args...))
}

func (l *Ledger) Info(args ...interface{}) {
	l.Write(InfoLevel, args...)
}

func (l *Ledger) Infof(f string, args ...interface{}) {
	l.Write(InfoLevel, fmt.Sprintf(f, args...))
}

func (l *Ledger) Warn(args ...interface{}) {
	l.Write(WarnLevel, args...)
}

func (l *Ledger) Warnf(f string, args ...interface{}) {
	l.Write(WarnLevel, fmt.Sprintf(f, args...))
}

func (l *Ledger) Error(args ...interface{}) {
	l.Write(ErrorLevel, args...)
}

func (l *Ledger) Errorf(f string, args ...interface{}) {
	l.Write(ErrorLevel, fmt.Sprintf(f, args...))
}

func (l *Ledger) Fatal(args ...interface{}) {
	l.Write(FatalLevel, args...)
}

func (l *Ledger) Fatalf(f string, args ...interface{}) {
	l.Write(FatalLevel, fmt.Sprintf(f, args...))
}
