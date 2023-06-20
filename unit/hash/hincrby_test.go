package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hincrby Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k, f, v := "hincrbykey", "field", "1"
		Expect(c.HSet(ctx, k, f, v).Err()).NotTo(HaveOccurred())
		Expect(c.HIncrBy(ctx, k, f, 1).Val()).To(Equal(int64(2)))
		Expect(c.HIncrBy(ctx, k, f, -1).Val()).To(Equal(int64(1)))
	})

	It("ERR hash value is not an integer", func() {
		k, f, v := "hincrbykey", "field", "v1"
		Expect(c.HSet(ctx, k, f, v).Err()).NotTo(HaveOccurred())
		Expect(c.HIncrBy(ctx, k, f, 1).Err().Error()).To(ContainSubstring("ERR hash value is not an integer"))
	})
})
