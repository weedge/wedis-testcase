package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hmget Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k1 := "hmgetk1"
		Expect(c.HSet(ctx, k1, "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, k1, "f2", "2").Err()).NotTo(HaveOccurred())
		Expect(c.HMGet(ctx, k1, "f1", "f2").Val()).To(Equal([]interface{}{"1", "2"}))
	})

	It("ERR wrong number of arguments", func() {
		k1 := "hmgetk2"
		Expect(c.HSet(ctx, k1, "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, k1, "f2", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HMGet(ctx, k1).Err().Error()).To(ContainSubstring("ERR wrong number of arguments"))
	})

	It("nil", func() {
		k1 := "hmgetk3"
		Expect(c.HSet(ctx, k1, "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, k1, "f2", "2").Err()).NotTo(HaveOccurred())
		Expect(c.HMGet(ctx, k1, "f1", "f3", "f2", "f4").Val()).To(Equal([]interface{}{"1", nil, "2", nil}))
	})
})
