package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hvals Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.HSet(ctx, "k1", "f1", "v1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, "k1", "f2", "v2").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, "k1", "f3", "v3").Err()).NotTo(HaveOccurred())
		Expect(c.HVals(ctx, "k1").Val()).To(Equal([]string{"v1", "v2", "v3"}))
	})

	It("empty", func() {
		Expect(c.HVals(ctx, "k2").Val()).To(Equal([]string{}))
	})
})
