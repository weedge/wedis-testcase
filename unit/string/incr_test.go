package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("INCR Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		key, val := "INCR", "1"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Incr(ctx, key).Val()).To(Equal(int64(2)))
		Expect(c.Incr(ctx, key).Val()).To(Equal(int64(3)))
	})

	It("ERR value is not an integer or out of range", func() {
		key, val := "INCR1", "v1"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Incr(ctx, key).Err().Error()).To(ContainSubstring("ERR value is not an integer or out of range"))
		Expect(c.Incr(ctx, key).Val()).To(Equal(int64(0)))
	})
	It("nokey", func() {
		Expect(c.Del(ctx, "nokey").Err()).NotTo(HaveOccurred())
		Expect(c.Incr(ctx, "nokey").Val()).To(Equal(int64(1)))
		Expect(c.Get(ctx, "nokey").Val()).To(Equal("1"))
	})
})
