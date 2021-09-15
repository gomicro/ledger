package ledger

import (
	"fmt"
	"testing"

	"github.com/gomicro/penname"

	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
)

func TestDefaultLogger(t *testing.T) {
	g := goblin.Goblin(t)
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
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", DebugLevel, msg)))
			mw.Reset()
			Info(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", InfoLevel, msg)))
			mw.Reset()
			Warn(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", WarnLevel, msg)))
			mw.Reset()
			Error(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", ErrorLevel, msg)))
			mw.Reset()
			Fatal(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", FatalLevel, msg)))
		})

		g.It("should log formatted levels", func() {
			msg := "time and relative dimension in space"

			Debugf("Additional Info: %v", msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [Additional Info: %v]\n", DebugLevel, msg)))
			mw.Reset()
			Infof("Additional Info: %v", msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [Additional Info: %v]\n", InfoLevel, msg)))
			mw.Reset()
			Warnf("Additional Info: %v", msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [Additional Info: %v]\n", WarnLevel, msg)))
			mw.Reset()
			Errorf("Additional Info: %v", msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [Additional Info: %v]\n", ErrorLevel, msg)))
			mw.Reset()
			Fatalf("Additional Info: %v", msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [Additional Info: %v]\n", FatalLevel, msg)))
		})

		g.It("should allow setting a threshold", func() {
			msg := "stormageddon"

			Threshold(DebugLevel)
			Debug(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", DebugLevel, msg)))
			mw.Reset()
			Info(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", InfoLevel, msg)))
			mw.Reset()
			Warn(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", WarnLevel, msg)))
			mw.Reset()
			Error(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", ErrorLevel, msg)))
			mw.Reset()
			Fatal(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", FatalLevel, msg)))

			Threshold(InfoLevel)
			Debug(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", FatalLevel, msg)))
			mw.Reset()
			Info(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", InfoLevel, msg)))
			mw.Reset()
			Warn(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", WarnLevel, msg)))
			mw.Reset()
			Error(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", ErrorLevel, msg)))
			mw.Reset()
			Fatal(msg)
			Expect(string(mw.Written())).To(Equal(fmt.Sprintf("%s: [%v]\n", FatalLevel, msg)))
		})
	})
}
