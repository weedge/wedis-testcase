package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("rpoplpush Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.Do(ctx, "LMCLEAR", "ll1", "ll2").Err()).NotTo(HaveOccurred())
		Expect(c.RPush(ctx, "ll1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.RPopLPush(ctx, "ll1", "ll2").Val()).To(Equal("v3"))
		Expect(c.LRange(ctx, "ll1", 0, -1).Val()).To(Equal([]string{"v1", "v2"}))
		Expect(c.LRange(ctx, "ll2", 0, -1).Val()).To(Equal([]string{"v3"}))
	})

	It("no exists", func() {
		Expect(c.Do(ctx, "LMCLEAR", "ll1", "ll2").Err()).NotTo(HaveOccurred())
		Expect(c.RPopLPush(ctx, "ll1", "ll2").Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
