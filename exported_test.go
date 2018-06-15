package ledger

import (
	"fmt"
	"testing"

	"github.com/gomicro/penname"

	. "github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestDefaultLogger(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Default Logging", func() {
		var mw *penname.PenName

		g.BeforeEach(func() {
			mw = penname.New()
			std.writer = mw
		})

		g.It("should log levels", func() {
			msg := "bad wolf"

			Debug(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]\n", DebugLevel, msg)))
			Info(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]\n", InfoLevel, msg)))
			Warn(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]\n", WarnLevel, msg)))
			Error(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]\n", ErrorLevel, msg)))
			Fatal(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]\n", FatalLevel, msg)))
		})

		g.It("should log formatted levels", func() {
			msg := "time and relative dimension in space"

			Debugf("Additional Info: %v", msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[Additional Info: %v\\]", DebugLevel, msg)))
			Infof("Additional Info: %v", msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[Additional Info: %v\\]", InfoLevel, msg)))
			Warnf("Additional Info: %v", msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[Additional Info: %v\\]", WarnLevel, msg)))
			Errorf("Additional Info: %v", msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[Additional Info: %v\\]", ErrorLevel, msg)))
			Fatalf("Additional Info: %v", msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[Additional Info: %v\\]", FatalLevel, msg)))
		})

		g.It("should allow setting a threshold", func() {
			msg := "stormageddon"

			Threshold(DebugLevel)
			Debug(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", DebugLevel, msg)))
			Info(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", InfoLevel, msg)))
			Warn(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", WarnLevel, msg)))
			Error(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", ErrorLevel, msg)))
			Fatal(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", FatalLevel, msg)))

			Threshold(InfoLevel)
			Debug(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", FatalLevel, msg)))
			Info(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", InfoLevel, msg)))
			Warn(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", WarnLevel, msg)))
			Error(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", ErrorLevel, msg)))
			Fatal(msg)
			Expect(string(mw.Written)).To(MatchRegexp(fmt.Sprintf("%s: .*exported_test.go:\\d+ \\[%v\\]", FatalLevel, msg)))
		})
	})
}
