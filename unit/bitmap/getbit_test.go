package bitmap

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GETBIT Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, val := "GETBITkey", "\x00\xff"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.GetBit(ctx,key,0).Val()).To(Equal(int64(0)))
		Expect(c.GetBit(ctx,key,8).Val()).To(Equal(int64(1)))
		Expect(c.GetBit(ctx, key, 100).Val()).To(Equal(int64(0)))	
	})
})
