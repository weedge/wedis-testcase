package list

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("lindex Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k := "lindexl1"
		Expect(c.RPush(ctx, k, "v1", "v2", "v3").Val()).To(Equal(int64(3)))
		Expect(c.LIndex(ctx, k, 0).Val()).To(Equal("v1"))
		Expect(c.LIndex(ctx, k, 1).Val()).To(Equal("v2"))
		Expect(c.LIndex(ctx, k, 2).Val()).To(Equal("v3"))
	})

	It("no key", func() {
		Expect(c.LIndex(ctx, "lindexnokey", 0).Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
