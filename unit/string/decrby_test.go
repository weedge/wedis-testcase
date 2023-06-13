package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DECRBY Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.Del(ctx, "k1").Err()).NotTo(HaveOccurred())
		Expect(c.DecrBy(ctx, "k1", 1).Val()).To(Equal(int64(-1)))
		Expect(c.DecrBy(ctx, "k1", 2).Val()).To(Equal(int64(-3)))
		Expect(c.DecrBy(ctx, "k1", -2).Val()).To(Equal(int64(-1)))

		Expect(c.Set(ctx, "k2", "v2", 0).Err()).NotTo(HaveOccurred())
		Expect(c.DecrBy(ctx, "k2", 1).Err().Error()).To(ContainSubstring("value is not an integer or out of range"))

		Expect(c.Set(ctx, "k3", "2", 0).Err()).NotTo(HaveOccurred())
		Expect(c.DecrBy(ctx, "k3", 1).Val()).To(Equal(int64(1)))
	})
})
