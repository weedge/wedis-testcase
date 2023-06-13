package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("del Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.Set(ctx, "k2", "v2", 0).Err()).NotTo(HaveOccurred())
		Expect(c.Del(ctx, "k2").Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, "k2").Err().Error()).To(ContainSubstring("redis: nil"))
	})
})
