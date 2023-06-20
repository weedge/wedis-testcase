package bitmap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SETBIT Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, val := "SETBITkey", "\x00\xff"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.GetBit(ctx, key, 0).Val()).To(Equal(int64(0)))
		Expect(c.SetBit(ctx, key, 0, 1).Err()).NotTo(HaveOccurred())
		Expect(c.GetBit(ctx, key, 0).Val()).To(Equal(int64(1)))
	})
})
