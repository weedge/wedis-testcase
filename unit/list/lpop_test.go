package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lpop Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "lpop"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.RPush(ctx, "lpopl1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.LPop(ctx, "lpopl1").Val()).To(Equal("v1"))
		Expect(c.LPop(ctx, "lpopl1").Val()).To(Equal("v2"))
		Expect(c.LPop(ctx, "lpopl1").Val()).To(Equal("v3"))
		Expect(c.LPop(ctx, "lpopl1").Val()).To(Equal(""))
	})

	It("no key", func() {
		k := "lpopnokey"
		Expect(c.Do(ctx, "LMCLEAR", k).Err()).NotTo(HaveOccurred())
		Expect(c.LPop(ctx, k).Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
