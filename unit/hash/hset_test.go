package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hset Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k, f, v := "key", "field", "val"
		Expect(c.HSet(ctx, k, f, v).Err()).NotTo(HaveOccurred())
		Expect(c.HGet(ctx, k, f).Val()).To(Equal(v))
		Expect(c.HGet(ctx, k, "other").Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
