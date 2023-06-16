package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hexists Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		k, f, v := "key", "field", "val"
		Expect(c.HSet(ctx, k, f, v).Err()).NotTo(HaveOccurred())
		Expect(c.HExists(ctx, k, f).Val()).To(BeTrue())
	})
	It("no", func() {
		Expect(c.HExists(ctx, "k1", "f11").Val()).To(BeFalse())
	})
})
