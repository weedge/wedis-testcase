package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DECR Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.Del(ctx, "k1", "k2", "k3").Err()).NotTo(HaveOccurred())
		Expect(c.Decr(ctx, "k1").Val()).To(Equal(int64(-1)))

		Expect(c.Set(ctx, "k2", "v2", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Decr(ctx, "k2").Err().Error()).To(ContainSubstring("value is not an integer or out of range"))

		Expect(c.Set(ctx, "k3", "2", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Decr(ctx, "k3").Val()).To(Equal(int64(1)))
	})
})
