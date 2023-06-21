package list

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lttl Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})
	It("ok", func() {
		k := "lttlkey"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.LPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "LEXPIRE", k, "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(100)))
		time.Sleep(1 * time.Second)
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(99)))
		Expect(c.Do(ctx, "LPERSIST", k).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(-1)))
	})

	It("no key", func() {
		k := "lttlnokey"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "LTTL", k).Val()).To(Equal(int64(-2)))
	})
})
