package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hmclear Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.HSet(ctx, "k1", "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, "k1", "f2", "1").Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "HMCLEAR", "k1").Val()).To(Equal(int64(1)))
		Expect(c.HGet(ctx, "k1", "f1").Err().Error()).To(ContainSubstring("redis: nil"))
	})

	It("hmclear empty", func() {
		Expect(c.Do(ctx, "HMCLEAR", "k1").Val()).To(Equal(int64(0)))
	})
})
