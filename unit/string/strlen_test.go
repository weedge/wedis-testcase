package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("STRLEN Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("strlen", func() {
		Expect(c.Set(ctx, "k1", -123, 0).Err()).NotTo(HaveOccurred())
		Expect(c.StrLen(ctx, "k1").Val()).To(Equal(int64(4)))
		Expect(c.Set(ctx, "k2", "1234567890", 0).Err()).NotTo(HaveOccurred())
		Expect(c.StrLen(ctx, "k2").Val()).To(Equal(int64(10)))
	})

	It("non-existing key", func() {
		Expect(c.StrLen(ctx, "k110").Val()).To(Equal(int64(0)))
	})
})
