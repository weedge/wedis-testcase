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
		Expect(c.Set(ctx, "STRLENk1", -123, 0).Err()).NotTo(HaveOccurred())
		Expect(c.StrLen(ctx, "STRLENk1").Val()).To(Equal(int64(4)))
		Expect(c.Set(ctx, "STRLENk2", "1234567890", 0).Err()).NotTo(HaveOccurred())
		Expect(c.StrLen(ctx, "STRLENk2").Val()).To(Equal(int64(10)))
	})

	It("non-existing key", func() {
		key := "STRLENk110"
		Expect(c.Del(ctx, key).Err()).NotTo(HaveOccurred())
		Expect(c.StrLen(ctx, key).Val()).To(Equal(int64(0)))
	})
})
