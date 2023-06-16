package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("hlen Cmd", func() {
	BeforeEach(func() {
	})

	AfterEach(func() {
	})

	It("ok", func() {
		Expect(c.HSet(ctx, "k1", "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, "k1", "f2", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HLen(ctx, "k1").Val()).To(Equal(int64(2)))
	})
})
