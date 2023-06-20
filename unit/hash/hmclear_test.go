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
		k := "hmclear"
		Expect(c.HSet(ctx, k, "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, k, "f2", "1").Err()).NotTo(HaveOccurred())
		Expect(c.Do(ctx, "HMCLEAR", k).Val()).To(Equal(int64(1)))
		Expect(c.HGet(ctx, k, "f1").Err().Error()).To(ContainSubstring("redis: nil"))
	})

	It("hmclear empty", func() {
		k := "hmclearnokey"
		Expect(c.Do(ctx, "HMCLEAR", k).Val()).To(Equal(int64(0)))
	})
})
