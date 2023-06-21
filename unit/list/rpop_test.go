package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("rpop Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "rpopkey"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.LPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.RPop(ctx, k).Val()).To(Equal("v1"))
		Expect(c.RPop(ctx, k).Val()).To(Equal("v2"))
		Expect(c.RPop(ctx, k).Val()).To(Equal("v3"))
		Expect(c.RPop(ctx, k).Val()).To(Equal(""))
		Expect(c.RPop(ctx, k).Val()).To(Equal(""))
	})

	It("no key", func() {
		k := "rpopnokey"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.RPop(ctx, k).Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
