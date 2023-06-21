package list

import (
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/weedge/wedis-testcase/help"
)

var _ = Describe("brpoplpush Cmd", func() {
	BeforeEach(func() {
		switch os.Getenv(help.EnvConfTestCase) {
		case "xdis-tikv":
			Expect(c.BRPopLPush(ctx, "brpoplpushl1", "brpoplpushl2", 1*time.Second).Err().Error()).To(ContainSubstring("ERR cmd not supported"))
			Skip(os.Getenv(help.EnvConfTestCase) + "case cmd unsupoort, skip test")
		}
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.RPush(ctx, "brpoplpushl1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.BRPopLPush(ctx, "brpoplpushl1", "brpoplpushl2", 1*time.Second).Val()).To(Equal("v3"))
		Expect(c.LRange(ctx, "brpoplpushl1", 0, -1).Val()).To(Equal([]string{"v1", "v2"}))
		Expect(c.LRange(ctx, "brpoplpushl2", 0, -1).Val()).To(Equal([]string{"v3"}))
	})

	It("no exists", func() {
		Expect(c.BRPopLPush(ctx, "brpoplpushll1", "ll2", 1*time.Second).Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
