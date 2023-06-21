package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("llen Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "llenl1"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.RPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.LLen(ctx, k).Val()).To(Equal(int64(3)))
	})

	It("no", func() {
		k := "llennokey"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.LLen(ctx, k).Val()).To(Equal(int64(0)))
	})
})
