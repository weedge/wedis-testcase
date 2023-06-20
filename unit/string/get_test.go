package string

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("get nokey", func() {
		key := "getnokey"
		Expect(c.Get(ctx, key).Err().Error()).To(ContainSubstring("redis: nil"))
	})

	It("set get", func() {
		key, val := "key", "val"
		Expect(c.Set(ctx, key, val, 0).Err()).NotTo(HaveOccurred())
		Expect(c.Get(ctx, key).Val()).To(Equal(val))
	})
})
