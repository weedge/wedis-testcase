package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("getrange Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, val := "getrange", "val"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.GetRange(ctx, key, 0, -1).Val()).To(Equal(val))
		Expect(c.GetRange(ctx, key, 0, 0).Val()).To(Equal("v"))
		Expect(c.GetRange(ctx, key, 0, 1).Val()).To(Equal("va"))
		Expect(c.GetRange(ctx, key, 0, 2).Val()).To(Equal("val"))
		Expect(c.GetRange(ctx, key, 0, -2).Val()).To(Equal("va"))
		Expect(c.GetRange(ctx, key, 0, -3).Val()).To(Equal("v"))
		Expect(c.GetRange(ctx, key, 0, -4).Val()).To(Equal("v"))
		Expect(c.GetRange(ctx, key, 0, -5).Val()).To(Equal("v"))
		Expect(c.GetRange(ctx, key, 0, -100).Val()).To(Equal("v"))

		Expect(c.GetRange(ctx, key, -1, 0).Val()).To(Equal(""))
		Expect(c.GetRange(ctx, key, -2, 0).Val()).To(Equal(""))
		Expect(c.GetRange(ctx, key, -3, 0).Val()).To(Equal("v"))
		Expect(c.GetRange(ctx, key, -4, 0).Val()).To(Equal("v"))
		Expect(c.GetRange(ctx, key, -5, 0).Val()).To(Equal("v"))
		Expect(c.GetRange(ctx, key, -5, 1).Val()).To(Equal("va"))
		Expect(c.GetRange(ctx, key, -5, 2).Val()).To(Equal("val"))
		Expect(c.GetRange(ctx, key, -5, 100).Val()).To(Equal("val"))
	})
})
