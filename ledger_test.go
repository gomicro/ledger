package ledger

import (
	"fmt"
	"testing"

	"github.com/gomicro/penname"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestLogging(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Logging", func() {
		g.It("should write log lines", func() {
			mw := penname.New()
			l := &Ledger{mw, DebugLevel}
			msg := "some sorta message"

			for lvl := FatalLevel; lvl <= DebugLevel; lvl++ {
				l.Write(lvl, msg)
				Expect(string(mw.Written)).To(Equal(fmt.Sprintf("%s: [%v]", lvl, msg)))
			}
		})

		g.It("should honor the threshold", func() {
			mw := penname.New()
			msg := "some sorta message"

			l := &Ledger{mw, DebugLevel}
			for lvl := FatalLevel; lvl <= DebugLevel; lvl++ {
				l.Write(lvl, msg)
				Expect(string(mw.Written)).To(Equal(fmt.Sprintf("%s: [%v]", lvl, msg)))
			}

			l = &Ledger{mw, InfoLevel}
			for lvl := FatalLevel; lvl <= InfoLevel; lvl++ {
				l.Write(lvl, msg)
				Expect(string(mw.Written)).To(Equal(fmt.Sprintf("%s: [%v]", lvl, msg)))
				mw.Clear()
			}
			l.Write(DebugLevel, msg)
			Expect(string(mw.Written)).NotTo(Equal(fmt.Sprintf("%s: [%v]", DebugLevel, msg)))
			Expect(string(mw.Written)).To(Equal(""))

			l = &Ledger{mw, WarnLevel}
			for lvl := FatalLevel; lvl <= WarnLevel; lvl++ {
				l.Write(lvl, msg)
				Expect(string(mw.Written)).To(Equal(fmt.Sprintf("%s: [%v]", lvl, msg)))
				mw.Clear()
			}
			for lvl := DebugLevel; lvl > WarnLevel; lvl-- {
				l.Write(lvl, msg)
				Expect(string(mw.Written)).NotTo(Equal(fmt.Sprintf("%s: [%v]", lvl, msg)))
				Expect(string(mw.Written)).To(Equal(""))
			}

			l = &Ledger{mw, ErrorLevel}
			for lvl := FatalLevel; lvl <= ErrorLevel; lvl++ {
				l.Write(lvl, msg)
				Expect(string(mw.Written)).To(Equal(fmt.Sprintf("%s: [%v]", lvl, msg)))
				mw.Clear()
			}
			for lvl := DebugLevel; lvl > ErrorLevel; lvl-- {
				l.Write(lvl, msg)
				Expect(string(mw.Written)).NotTo(Equal(fmt.Sprintf("%s: [%v]", lvl, msg)))
				Expect(string(mw.Written)).To(Equal(""))
			}

			l = &Ledger{mw, FatalLevel}
			l.Write(FatalLevel, msg)
			Expect(string(mw.Written)).To(Equal(fmt.Sprintf("%s: [%v]", FatalLevel, msg)))
			mw.Clear()
			for lvl := DebugLevel; lvl > FatalLevel; lvl-- {
				l.Write(lvl, msg)
				Expect(string(mw.Written)).NotTo(Equal(fmt.Sprintf("%s: [%v]", lvl, msg)))
				Expect(string(mw.Written)).To(Equal(""))
			}
		})
	})
}
