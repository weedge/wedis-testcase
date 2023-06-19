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
		Expect(c.HSet(ctx, "hlen", "f1", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HSet(ctx, "hlen", "f2", "1").Err()).NotTo(HaveOccurred())
		Expect(c.HLen(ctx, "hlen").Val()).To(Equal(int64(2)))
	})
})
