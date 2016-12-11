package ledger

import (
	"fmt"
	"io"
	"sync"
)

type Ledger struct {
	writer    io.Writer
	threshold Level
	mu        sync.Mutex
}

func New(w io.Writer, l Level) *Ledger {
	return &Ledger{
		writer:    w,
		threshold: l,
	}
}

func (l *Ledger) Debug(args ...interface{}) {
	l.write(DebugLevel, args...)
}

func (l *Ledger) Debugf(f string, args ...interface{}) {
	l.write(DebugLevel, fmt.Sprintf(f, args...))
}

func (l *Ledger) Info(args ...interface{}) {
	l.write(InfoLevel, args...)
}

func (l *Ledger) Infof(f string, args ...interface{}) {
	l.write(InfoLevel, fmt.Sprintf(f, args...))
}

func (l *Ledger) Warn(args ...interface{}) {
	l.write(WarnLevel, args...)
}

func (l *Ledger) Warnf(f string, args ...interface{}) {
	l.write(WarnLevel, fmt.Sprintf(f, args...))
}

func (l *Ledger) Error(args ...interface{}) {
	l.write(ErrorLevel, args...)
}

func (l *Ledger) Errorf(f string, args ...interface{}) {
	l.write(ErrorLevel, fmt.Sprintf(f, args...))
}

func (l *Ledger) Fatal(args ...interface{}) {
	l.write(FatalLevel, args...)
}

func (l *Ledger) Fatalf(f string, args ...interface{}) {
	l.write(FatalLevel, fmt.Sprintf(f, args...))
}

func (l *Ledger) write(lvl Level, args ...interface{}) {
	if lvl <= l.threshold {
		out, args := fmt.Sprintf("[%v]", args[0]), args[1:]
		for _, arg := range args {
			out += fmt.Sprintf(" [%v]", arg)
		}

		l.mu.Lock()
		defer l.mu.Unlock()
		fmt.Fprintf(l.writer, "%s: %v\n", lvl, out)
	}
}

func (l *Ledger) Threshold(level Level) {
	l.threshold = level
}
