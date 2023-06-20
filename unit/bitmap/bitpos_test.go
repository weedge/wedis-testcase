package bitmap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BITPOS Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, val := "BITPOSkey", "\xff\xf0\x00"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.BitPos(ctx, key, 0).Val()).To(Equal(int64(12)))
		val = "\x00\xff\xf0"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.BitPos(ctx, key, 0).Val()).To(Equal(int64(0)))
		Expect(c.BitPos(ctx, key, 1, 0).Val()).To(Equal(int64(8)))
		Expect(c.BitPos(ctx, key, 1, 1).Val()).To(Equal(int64(8)))
		Expect(c.BitPos(ctx, key, 1, 2).Val()).To(Equal(int64(16)))
		Expect(c.BitPos(ctx, key, 1, 3).Val()).To(Equal(int64(-1)))
	})
})
