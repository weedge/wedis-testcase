package set

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sttl Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "sttl"
		Expect(c.SAdd(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "sEXPIRE", k, "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "sTTL", k).Val()).To(Equal(int64(100)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "sTTL", k).Val()).To(Equal(int64(99)))
		Expect(c.Do(ctx, "sPERSIST", k).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "sTTL", k).Val()).To(Equal(int64(-1)))
	})

	It("no key", func() {
		Expect(c.Do(ctx, "sTTL", "nokeysttl").Val()).To(Equal(int64(-2)))
	})
})
