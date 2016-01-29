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
