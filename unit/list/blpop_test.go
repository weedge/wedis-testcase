package list

import (
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/weedge/wedis-testcase/help"
)

var _ = Describe("blpop Cmd", func() {
	BeforeEach(func() {
		switch os.Getenv(help.EnvConfTestCase) {
		case "xdis-tikv":
			Expect(c.BLPop(ctx, 1*time.Second, "blpopl1").Err().Error()).To(ContainSubstring("ERR cmd not supported"))
			Skip(os.Getenv(help.EnvConfTestCase) + "case cmd unsupoort, skip test")
		}
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.RPush(ctx, "blpopl1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.BLPop(ctx, 1*time.Second, "blpopl1").Val()).To(Equal([]string{"blpopl1", "v1"}))
		Expect(c.BLPop(ctx, 1*time.Second, "blpopl1").Val()).To(Equal([]string{"blpopl1", "v2"}))
		Expect(c.BLPop(ctx, 1*time.Second, "blpopl1").Val()).To(Equal([]string{"blpopl1", "v3"}))
		Expect(c.BLPop(ctx, 1*time.Second, "blpopl1").Val()).To(Equal([]string{}))
	})

	It("block no exists", func() {
		s := time.Now()
		Expect(c.BLPop(ctx, 2*time.Second, "blpopl110").Val()).To(Equal([]string{}))
		Expect(time.Since(s).Seconds() > 2).To(BeTrue())
	})

})
