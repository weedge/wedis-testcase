package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("rpush Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.Do(ctx, "LMCLEAR", "ll1").Err()).NotTo(HaveOccurred())
		Expect(c.RPush(ctx, "ll1", "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.LPop(ctx, "ll1").Val()).To(Equal("v1"))
		Expect(c.LPop(ctx, "ll1").Val()).To(Equal("v2"))
		Expect(c.LPop(ctx, "ll1").Val()).To(Equal("v3"))
		Expect(c.LPop(ctx, "ll1").Val()).To(Equal(""))
	})
})
