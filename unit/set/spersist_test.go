package set

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("spersist Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})
	It("no exists key", func() {
		k := "spersistnokey"
		Expect(c.Do(ctx, "SMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "sPERSIST", k).Val()).To(Equal(int64(0)))
	})

	It("ok", func() {
		k := "spersist"
		Expect(c.Do(ctx, "SMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.SAdd(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.Do(ctx, "sEXPIRE", k, "100").Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "sTTL", k).Val()).To(Equal(int64(100)))
		Expect(c.Do(ctx, "sPERSIST", k).Val()).To(Equal(int64(1)))
		Expect(c.Do(ctx, "sTTL", k).Val()).To(Equal(int64(-1)))
	})
})
